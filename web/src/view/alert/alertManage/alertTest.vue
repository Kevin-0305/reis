<template>
  <div>
    <div class="gva-search-box">
        <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
          <el-form-item>
          </el-form-item>
          <!-- <el-form-item label="集群名称">
            <el-input v-model="" placeholder="搜索条件" />
          </el-form-item> -->
        </el-form>
      </div>
      <div class="gva-table-box">
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        >
        <el-table-column align="left" label="ID" prop="id" width="500" />
        <el-table-column align="left" label="名称" prop="title" width="500" />
        <el-table-column align="left" label="按钮组">
            <template #default="scope">
            <el-button type="primary" icon="watch" size="small" class="table-button" @click="toTarget(scope.row)">告警测试</el-button>
            </template>
        </el-table-column>
        </el-table>
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
  getAlertTypes,
  alertTest,
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
const searchInfo = ref({})
const route = useRoute()


// 重置
const onReset = () => {
  searchInfo.value = {}
}

const router = useRouter()


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
  const table = await getAlertTypes({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    console.log(table.data)
    tableData.value = table.data.list
  }
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
</style>
