<template>
  <div>
    <div class="gva-search-box">
        <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
          <el-form-item>
            <el-select v-model="searchInfo.serviceName" placeholder="请选择prometheus服务">
              <el-option
                v-for="item in promServices"
                :key="item"
                :label="item"
                :value="item">
              </el-option>
            </el-select>
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
        <el-table-column align="left" label="cpu" prop="cpu" width="200">
          <template #default="scope">{{cpuFilter(scope.row.cpu)}}</template>
        </el-table-column>
                <el-table-column align="left" label="cpu" prop="cpu" width="200">
          <template #default="scope">{{gpuFilter(scope.row.gpu)}}</template>
        </el-table-column>
        <el-table-column align="left" label="disk(T)" prop="disk" width="120" />
        <el-table-column align="left" label="instance" prop="instance" width="150" />
        <el-table-column align="left" label="ip" prop="ip" width="120" />
        <el-table-column align="left" label="job" prop="job" width="200" />
        <el-table-column align="left" label="mac" prop="mac" width="150" />
        <el-table-column align="left" label="memory(G)" prop="memory" width="120" />
        <el-table-column align="left" label="os" prop="os" width="120" />
        <el-table-column align="left" label="user" prop="user" width="120" />
        <el-table-column align="left" label="按钮组">
            <template #default="scope">
            <el-button type="primary" icon="watch" size="small" class="table-button" @click="toTarget(scope.row)">实时数据</el-button>
            </template>
        </el-table-column>
        </el-table>
      </div>
  </div>
</template>

<script>
export default {
  name: 'EsCluter',
}
</script>

<script setup>
import {
  getPromServiceList,
  getActiveHostList,
} from '@/api/promHost'


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
const promServices = ref([])
const promServiceSelect = ref("")
const route = useRoute()


// 重置
const onReset = () => {
  searchInfo.value = {}
}

// 搜索
const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  if (searchInfo.value.monitor === ""){
      searchInfo.value.monitor=null
  }
  if (searchInfo.value.TLS === ""){
      searchInfo.value.TLS=null
  }
  if (searchInfo.value.auth === ""){
      searchInfo.value.auth=null
  }
  getTableData()
}

const router = useRouter()

const toTarget = (row) => {
  console.log("hostRow",row.instance)
  router.push({name:"hostRealInfo",params:{instance:row.instance,cpu:row.cpu,gpu:row.gpu,memory:row.memory,disk:row.disk,serviceName:promServiceSelect.value}})
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getActiveHostList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    console.log(table.data)
    tableData.value = table.data.list
  }
}

const loadPromServiceList = async() => {
  const table = await getPromServiceList()
  if (table.code === 0) {
    promServices.value = table.data.list
  }
}

//cpu名字简化过滤
const cpuFilter = (value) => {
  if (value.indexOf("AMD") !== -1) {
    const stringList = value.split(" ")
    return stringList[3]
  } else if (value.indexOf("Intel") !== -1) {
    const stringList = value.split(" ")
    for (let i = 0; i < stringList.length; i++) {
      if (stringList[i].indexOf("-") !== -1) {
        return stringList[i]
      }
    }
  } else {
    return value
  }
}

//gpu名字过滤简化
const gpuFilter = (value) => {
  if (value.indexOf("GeForce") !== -1) {
    return value.replace("GeForce ", "").replace("RTX ", "").replace("GTX ", "")
  } else if (value.indexOf("AMD") !== -1) {
    return "AMD显卡"
  } else {
    return "无显卡"
  }
}


const initPage = async() => {
  await loadPromServiceList()
  if ( promServices.value.length > 0) {
  promServiceSelect.value = promServices.value[0]
  searchInfo.value.serviceName = promServiceSelect
  getTableData()
  }else{
  ElMessageBox.alert('暂无可用的Prometheus服务,请联系管理员！')
  }
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
