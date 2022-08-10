<template>
  <div>
  <div class="gva-search-box">
        <el-form ref="searchForm" :inline="true" :model="searchInfo">
          <el-form-item label="instance">
            <el-input v-model="searchInfo.instance" placeholder="instance" />
          </el-form-item>
          <el-form-item label="job">
            <el-input v-model="searchInfo.job" placeholder="job" />
          </el-form-item>
          <el-form-item label="告警名称">
            <el-input v-model="searchInfo.alertName" placeholder="告警名称" />
          </el-form-item>

          <el-form-item label="severity">
            <el-select v-model="searchInfo.severity" clearable placeholder="请选择">
              <el-option
                v-for="item in severityOptions"
                :key="item.value"
                :label="item.label"
                :value="item.value"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="状态">
            <el-select v-model="searchInfo.status" clearable placeholder="请选择">
              <el-option
                v-for="item in statusOptions"
                :key="item.value"
                :label="item.label"
                :value="item.value"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="告警时间">
            <!-- <div class="block"> -->
              <!-- <span class="demonstration">With shortcuts</span> -->
              <el-date-picker
                v-model="timeRage"
                type="datetimerange"
                :shortcuts="shortcuts"
                range-separator="到"
                start-placeholder="开始时间"
                end-placeholder="结束时间"
              />
            <!-- </div> -->
          </el-form-item>

          <el-form-item>
            <el-button size="small" type="primary" icon="search" @click="onSubmit">查询</el-button>
            <el-button size="small" icon="refresh" @click="onReset">重置</el-button>
          </el-form-item>
        </el-form>
      </div>
      <div class="gva-table-box">
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        >
        <el-table-column type="expand">
          <template #default="scope">
            <el-form label-position="left" inline class="demo-table-expand">
              <el-form-item label="说明">
                <span>{{scope.row.annotations.description}}</span>
              </el-form-item>
              <el-form-item label="概要">
                <span>{{scope.row.annotations.summary}}</span>
              </el-form-item>
              <el-form-item label="开始时间">
                <span>{{getTime(scope.row.startsAt)}}</span>
              </el-form-item>
              <el-form-item label="结束时间">
                <span>{{getTime(scope.row.endsAt)}}</span>
              </el-form-item>
            </el-form>
          </template>
        </el-table-column>
        <el-table-column align="left" label="状态" prop="status" width="200" />
        <el-table-column align="left" label="告警名称" prop="labels.alertname" width="200" />
        <el-table-column align="left" label="job" prop="labels.job" width="200" />
        <el-table-column align="left" label="instance" prop="labels.instance" width="200" />
        <el-table-column align="left" label="severity" prop="labels.severity" width="200" />


        <el-table-column align="left" label="按钮组">
            <template>
            <el-button type="primary" icon="watch" size="small" class="table-button">查看详细报告</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
        <el-pagination
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          layout="total, sizes, prev, pager, next, jumper"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
      </div>
  </div>
</template>

<script>
export default {
  name: 'alertTest',
}
</script>

<script setup>
import {
  getPrometheusAlertHistory,
} from '@/api/alert'


// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { nextTick, ref, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'



// 自动化生成的字典（可能为空）以及字段


// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const timeRage = ref('')
const searchInfo = ref({
  "alertName":"",
  "status":"",
  "Job":"",
  "instance":"",
  "severity":""
})
const route = useRoute()
const statusOptions = [
  {
    value: '',
    label: '全部',
  },
  {
    value: 'firing',
    label: 'firing',
  },
  {
    value: 'resolved',
    label: 'resolved',
  },
]

const severityOptions = [
  {
    value: '',
    label: '全部',
  },
  {
    value: 'error',
    label: 'error',
  },
  {
    value: 'warning',
    label: 'warning',
  },
]

const shortcuts = [
  {
    text: '昨天',
    value: () => {
      const end = new Date().setHours(0, 0, 0, 0)
      const start = new Date()
      start.setTime(end - 3600 * 1000 * 24 * 1)
      return [start, end]
    },
  },
  {
    text: '一周内',
    value: () => {
      const end = new Date()
      const start = new Date()
      start.setTime(start.getTime() - 3600 * 1000 * 24 * 7)
      return [start, end]
    },
  },
  {
    text: '一月内',
    value: () => {
      const end = new Date()
      const start = new Date()
      start.setTime(start.getTime() - 3600 * 1000 * 24 * 30)
      return [start, end]
    },
  },
  {
    text: '三月内',
    value: () => {
      const end = new Date()
      const start = new Date()
      start.setTime(start.getTime() - 3600 * 1000 * 24 * 90)
      return [start, end]
    },
  },
]


// 重置
const onReset = () => {
  searchInfo.value = {
    "alertName":"",
    "status":"",
    "Job":"",
    "instance":"",
    "severity":""
  }
  timeRage.value = ['', '']
}

const router = useRouter()

//时区转换
const getTime = (time) => {
  return formatDate(new Date(time).getTime(), 'yyyy-MM-dd hh:mm:ss')
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

//测试
const toTarget = async(row) => {
  const result = await alertTest({"alertType":row.name})
  if(result.code === 0){
    ElMessageBox.alert(result.data, '提示')
  }
  console.log(row)
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  searchInfo.value.pageInfo = {
    "page": page.value,
    "pageSize": pageSize.value
  }
  if (timeRage.value !== ['', '']) {
    searchInfo.value.alertTimeGte = timeRage.value[0]
    searchInfo.value.alertTimeLte = timeRage.value[1]
  }
  const table = await getPrometheusAlertHistory(searchInfo.value)
  if (table.code === 0) {
    console.log(table.data)
    tableData.value = table.data.list
  }
}

const onSubmit = () => {
  getTableData()
}

const initPage = async() => {
  getTableData()
}

initPage()
//getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留


// 获取需要的字典 可能为空 按需保留
//setOptions()




</script>

<style>
  .demo-table-expand {
    font-size: 0;
  }
  .demo-table-expand label {
    width: 90px;
    color: #99a9bf;
  }
  .demo-table-expand .el-form-item {
    margin-right: 0;
    margin-bottom: 0;
    width: 50%;
  }
  .block {
    padding: 30px 0;
    text-align: center;
    border-right: solid 1px var(--el-border-color);
    flex: 1;
  }
  .block:last-child {
    border-right: none;
  }
  .block .demonstration {
    display: block;
    color: var(--el-text-color-secondary);
    font-size: 14px;
    margin-bottom: 20px;
  }
</style>
