<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="集群名称">
          <el-input v-model="searchInfo.cluterName" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="集群状态">
          <el-input v-model="searchInfo.status" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="集群ES版本">
          <el-input v-model="searchInfo.version" placeholder="搜索条件" />
        </el-form-item>
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
        <!-- <el-table-column align="left" label="日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column> -->
        <el-table-column align="left" label="集群名称" prop="cluterName" width="120" />
        <el-table-column align="left" label="集群状态" prop="status" width="120" >
          <template #default="scope">{{ formatStatus(scope.row.status) }}</template>
        </el-table-column>
        <el-table-column align="left" label="集群ES版本" prop="version" width="120" />
        <el-table-column align="left" label="集群地址" prop="address" width="120" />
        <el-table-column align="left" label="是否启用监控" prop="monitor" width="120">
            <template #default="scope">{{ formatBoolean(scope.row.monitor) }}</template>
        </el-table-column>
        <el-table-column align="left" label="TLS" prop="TLS" width="120">
            <template #default="scope">{{ formatBoolean(scope.row.TLS) }}</template>
        </el-table-column>
        <el-table-column align="left" label="身份验证" prop="auth" width="120">
            <template #default="scope">{{ formatBoolean(scope.row.auth) }}</template>
        </el-table-column>
        <!-- <el-table-column align="left" label="用户名" prop="userName" width="120" />
        <el-table-column align="left" label="密码" prop="password" width="120" /> -->
        <el-table-column align="left" label="描述" prop="description" width="360" />

        <el-table-column align="left" label="集群归属" min-width="200">
          <template #default="scope">
            <el-cascader
              v-model="scope.row.groupIds"
              :options="groupOptions"
              :show-all-levels="false"
              collapse-tags
              :props="{ multiple:true,checkStrictly: true,label:'groupName',value:'groupId',disabled:'disabled',emitPath:false}"
              :clearable="false"
              @visible-change="(flag)=>{changeGroup(scope.row,flag)}"
              @remove-tag="()=>{changeGroup(scope.row,false)}"
            />
          </template>
        </el-table-column>

        <el-table-column align="left" label="按钮组">
            <template #default="scope">
            <el-button type="text" icon="edit" size="small" class="table-button" @click="updateEsCluterFunc(scope.row)">变更</el-button>
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
        <el-form-item label="集群名称:">
          <el-input v-model="formData.cluterName" clearable placeholder="请输入" />
        </el-form-item>
        <!-- <el-form-item label="集群状态:">
          <el-input v-model.number="formData.status" clearable placeholder="请输入" />
        </el-form-item> -->
        <el-form-item label="集群ES版本:" v-if = "versionVisible" >
          <el-input v-model="formData.version"  clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="集群地址:">
          <el-input v-model="formData.address" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="是否启用监控:">
          <el-switch v-model="formData.monitor" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
        </el-form-item>
        <el-form-item label="TLS:">
          <el-switch v-model="formData.TLS" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
        </el-form-item>
        <el-form-item label="身份验证:">
          <el-switch v-model="formData.auth" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
        </el-form-item>
        <el-form-item label="用户名:" v-if = "formData.auth">
          <el-input v-model="formData.userName" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="密码:" v-if = "formData.auth">
          <el-input v-model="formData.password" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="描述:">
          <el-input v-model="formData.description" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="集群归属" prop="authorityId">
            <el-cascader
              v-model="formData.groupIds"
              style="width:100%"
              :options="groupOptions"
              :show-all-levels="false"
              :props="{ multiple:true,checkStrictly: true,label:'groupName',value:'groupId',disabled:'disabled',emitPath:false}"
              :clearable="false"
            />
          </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="checkEsCluterStatus">测 试</el-button>
          <el-button size="small" @click="closeDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'EsCluter',
}
</script>

<script setup>
import {
  createEsCluter,
  deleteEsCluter,
  deleteEsCluterByIds,
  updateEsCluter,
  findEsCluter,
  getEsCluterList,
  checkEsCluter,
  setEsCluterGroup,
} from '@/api/esCluter'

import {
  getProjectGroupTreeList
} from '@/api/esProjectGroup'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { nextTick, ref, watch } from 'vue'

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
        cluterName: '',
        status: 0,
        version: '',
        address: '',
        monitor: false,
        TLS: false,
        auth: false,
        userName: '',
        password: '',
        description: '',
        groupIds: [],
        })

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
const versionVisible = ref(false)

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
  const table = await getEsCluterList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

watch(tableData, () => {
  setGroupIds()
})
const setGroupIds = () => {
  tableData.value && tableData.value.forEach((cluter) => {
    const groupIds = cluter.group && cluter.group.map(i => {
      return i.ID
    })
    cluter.groupIds = groupIds
  })
}

const groupOptions = ref([])
const setOptions = async (groupData) =>{
    groupOptions.value = []
    setGroupOptions(groupData, groupOptions.value)
}

const setGroupOptions = (GroupData, optionsData) => {
  GroupData &&
        GroupData.forEach(item => {
          if (item.children && item.children.length) {
            const option = {
              groupId: item.ID,
              groupName: item.name,
              children: []
            }
            setGroupOptions(item.children, option.children)
            optionsData.push(option)
          } else {
            const option = {
              groupId: item.ID,
              groupName: item.name,
            }
            optionsData.push(option)
          }
        })
}

const initPage = async() => {
  getTableData()
  const res = await getProjectGroupTreeList({ page: 1, pageSize: 999 })
  setOptions(res.data.list)
}

initPage()
//getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留


// 获取需要的字典 可能为空 按需保留
//setOptions()


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
            deleteEsCluterFunc(row)
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
      const res = await deleteEsCluterByIds({ ids })
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
const updateEsCluterFunc = async(row) => {
    const res = await findEsCluter({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.reesc
        formData.value.groupIds = res.data.reesc.group && res.data.reesc.group.map(i => {
          return i.ID
        })
        versionVisible.value = true
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteEsCluterFunc = async (row) => {
    const res = await deleteEsCluter({ ID: row.ID })
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

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    versionVisible.value = false
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        cluterName: '',
        status: 0,
        version: '',
        address: '',
        monitor: false,
        TLS: false,
        auth: false,
        userName: '',
        password: '',
        description: '',
        }
}
// 弹窗确定
const enterDialog = async () => {
      let res
      formData.value.address = formData.value.address.replace(/\s+/g, "")
      switch (type.value) {
        case 'create':
          res = await createEsCluter(formData.value)
          break
        case 'update':
          res = await updateEsCluter(formData.value)
          break
        default:
          res = await createEsCluter(formData.value)
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

const checkEsCluterStatus = async () => {
  formData.value.address = formData.value.address.replace(/\s+/g, "")
  const res = await checkEsCluter(formData.value)
  console.log(res)
  if (res.code === 0) {
    if (res.data.status != 0 ){
      ElMessage({
        type: 'success',
        message: '集群连接成功'
      })
    }else{
      ElMessage({
        type: 'error',
        message: '集群连接失败'
      })
    }
  }
}

const changeGroup = async(row, flag) => {
  if (flag) {
    return
  }

  await nextTick()
  const res = await setEsCluterGroup({
    cluterId: row.ID,
    groupIds: row.groupIds
  })
  if (res.code === 0) {
    ElMessage({ type: 'success', message: '角色设置成功' })
  }
}


const formatStatus = (state) => {
  if (state === 0) {
    return '不可用'
  } else if (state === 1) {
    return '可用'
  }

}
</script>

<style>
</style>
