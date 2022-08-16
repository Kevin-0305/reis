<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button size="small" icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button size="small" type="primary" icon="plus" @click="openDialog">新增</el-button>
            <el-popover v-model:visible="deleteVisible" placement="top" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin-top: 8px;">
                <el-button size="small" type="text" @click="deleteVisible = false">取消</el-button>
                <el-button size="small" type="primary" @click="onDelete">确定</el-button>
            </div>
            <template #reference>
                <el-button icon="delete" size="small" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="deleteVisible = true">删除</el-button>
            </template>
            </el-popover>
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.timestamp)}}</template>
        </el-table-column>
        <el-table-column align="left" label="构建编号" prop="number" width="120" />
        <el-table-column align="left" label="构建名" prop="displayName" width="120" />
        <el-table-column align="left" label="状态" prop="building" width="120" >
          <template #default="scope">{{ scope.row.building ? '构建中' : '构建完成' }}</template>
        </el-table-column>  
        <el-table-column align="left" label="构建结果" prop="result" width="120" >
          <template #default="scope">
            <template v-if="scope.row.result == 'SUCCESS'">
              <el-tag size="small" type="success">成功</el-tag>
            </template>
            <template v-else-if="scope.row.result == 'FAILURE'">
              <el-tag size="small" type="danger">失败</el-tag>
            </template>
            <template v-else-if="scope.row.result == 'ABORTED'">
              <el-tag size="small" type="warning">中断</el-tag>
            </template>
            <template v-else-if="scope.row.result == 'UNSTABLE'">
              <el-tag size="small" type="primary">不稳定</el-tag>
            </template>
            <template v-else>
              <el-tag size="small" type="info">未知</el-tag>
            </template>
          </template>
        </el-table-column>
        <el-table-column align="left" label="按钮组">
            <template #default="scope">
            <el-button type="text" icon="edit" size="small" class="table-button" @click="getJenkinsBuildLog(scope.row)">构建日志</el-button>
            <el-button type="text" icon="edit" size="small" class="table-button" @click="updateJenkinsBuildHistoryFunc(scope.row)">灰度测试发布</el-button>
            <el-button type="text" icon="edit" size="small" class="table-button" @click="updateJenkinsBuildHistoryFunc(scope.row)">生产环境发布</el-button>
            <el-button type="text" icon="delete" size="small" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="构建编号:">
          <el-input v-model.number="formData.number" clearable placeholder="请输入" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>

    <el-drawer
    v-model="drawer"
    title="构建日志"
    :direction="direction"
    :before-close="handleClose"
    size="50%"
    >
    <div>{{buildLog}}</div>
    </el-drawer>
    <!-- <el-drawer v-model="drawer" :direction="direction">
      <template #title>
        <h4>构建日志</h4>
      </template>
      <template #default>
        <span>Hi, there!</span>
      </template>
  </el-drawer> -->
  </div>
</template>

<script>
export default {
  name: 'JenkinsBuildHistory'
}
</script>

<script setup>
import {
  createJenkinsBuildHistory,
  deleteJenkinsBuildHistory,
  deleteJenkinsBuildHistoryByIds,
  updateJenkinsBuildHistory,
  findJenkinsBuildHistory,
  getJenkinsBuildHistoryList,
  getJenkinsBuildLogText
} from '@/api/cicdJenkinsBuildHistory'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref } from 'vue'

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
        number: 0,
        })

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
const drawer = ref(false)
const direction = ref('rtl')
const buildLog = ref('')

// 重置
const onReset = () => {
  searchInfo.value = {}
}

// 搜索
const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  getTableData()
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
  searchInfo.value.serverId = 1
  searchInfo.value.jobName = "testproject1"
  const table = await getJenkinsBuildHistoryList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteJenkinsBuildHistoryFunc(row)
        })
    }


// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async() => {
      const ids = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          ids.push(item.ID)
        })
      const res = await deleteJenkinsBuildHistoryByIds({ ids })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === ids.length && page.value > 1) {
          page.value--
        }
        deleteVisible.value = false
        getTableData()
      }
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateJenkinsBuildHistoryFunc = async(row) => {
    const res = await findJenkinsBuildHistory({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.rejbh
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteJenkinsBuildHistoryFunc = async (row) => {
    const res = await deleteJenkinsBuildHistory({ ID: row.ID })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 查看构建日志
const getJenkinsBuildLog = async(row) => {
    const res = await getJenkinsBuildLogText({ buildNumber: 33,serverId:1,jobName:"testproject1" })
    drawer.value = true
    buildLog.value = res.data.logText.replace(/\n/g, '<br>')
    // if (res.code === 0) {
    //     ElMessageBox.alert(res.data, '构建日志')
    // }
}


// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        number: 0,
        }
}
// 弹窗确定
const enterDialog = async () => {
      let res
      switch (type.value) {
        case 'create':
          res = await createJenkinsBuildHistory(formData.value)
          break
        case 'update':
          res = await updateJenkinsBuildHistory(formData.value)
          break
        default:
          res = await createJenkinsBuildHistory(formData.value)
          break
      }
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '创建/更改成功'
        })
        closeDialog()
        getTableData()
      }
}
</script>

<style>
</style>
