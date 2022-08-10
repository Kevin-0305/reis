<template>
  <!-- <div class="chart-container">
    <chart height="100%" width="100%" ref="chart" />
  </div> -->
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="集群名称">
          <el-input  placeholder="搜索条件" />
        </el-form-item>
      </el-form>
    </div>
  </div> 
  
  <el-row :gutter="20">
  <el-col :span="4">
  <el-row>
    <img src="@/assets/CPU.png"  width="100" height="100" style="object-fit: contain" class="gva-top-card-right" alt/>
    <div>cpu型号:{{cpu}}</div>
  </el-row> 
  <el-row>
    <img src="@/assets/Memory.png"  width="100" height="100" style="object-fit: contain" class="gva-top-card-right" alt/>
    <div>内存大小:{{memory}}G</div>
  </el-row>
  </el-col>
  <el-col :span="4"><div class="grid-content bg-purple">
    <div>
    <img src="@/assets/GPU.png"  width="100" height="100" style="object-fit: contain;display:block" class="gva-top-card-right" alt/>
    <div>    显卡型号:{{gpu}}</div>
    </div>
  
    <div>
    <div>
    <img src="@/assets/Disk.png"  width="100" height="100" style="object-fit: contain;display:block" class="gva-top-card-right" alt/>
    </div>
    <div>硬盘大小:{{disk}}T</div>
    </div>
  </div></el-col>

  <el-col :span="8"><div class="grid-content bg-purple">
    <TemperatureGauge width="'100px'" height="'100px'" :title="'cpu温度'" ref="cpuTemChart"></TemperatureGauge>
  </div></el-col>
  <el-col :span="8"><div class="grid-content bg-purple">
    <TemperatureGauge width="'100px'" height="'100px'" :title="'显卡温度'" ref="gpuTemChart"></TemperatureGauge>
  </div></el-col>
</el-row>

<el-row :gutter="20">
  <el-col :span="8"><div class="grid-content bg-purple">
    <disk-bar style="height:750px" :title="'硬盘使用量'" ref="diskUseChart"></disk-bar>
  </div></el-col>
  <el-col :span="8"><div class="grid-content bg-purple">
    <el-raw>
      <data-line width="'200px'" height="'200px'" :title="'CPU使用率'" ref="cpuUseChart"></data-line>
    </el-raw>
    <el-raw>
      <data-line width="'200px'" height="'200px'" :title="'GPU使用率'" ref="gpuUseChart"></data-line>
    </el-raw>
  </div></el-col>
  <el-col :span="8"><div class="grid-content bg-purple">
    <el-raw>
      <data-line width="'200px'" height="'200px'" :title="'内存使用率'" ref="memoryUseChart"></data-line>
    </el-raw>
    <el-raw>
      <data-line width="'200px'" height="'200px'"></data-line>
    </el-raw>
  </div></el-col>
</el-row>

<el-row :gutter="20">
  <el-col :span="4"><div class="grid-content bg-purple"></div></el-col>
  <el-col :span="16"><div class="grid-content bg-purple"></div></el-col>
  <el-col :span="4"><div class="grid-content bg-purple"></div></el-col>
</el-row>
  <!-- <data-line width="'200px'" height="'200px'"></data-line> -->
</template>

<script>

export default {
  name: 'KeyboardChart',
  components: { Chart,DataLine,TemperatureGauge,DiskBar},

mounted(){
  // setInterval()
  }
}

</script>


<script setup>

import {ref} from 'vue'
import Chart from '@/components/echarts/Keyboard.vue'
import DataLine from "@/components/echarts/memoryLine.vue";
import TemperatureGauge from "@/components/echarts/temperatureGauge.vue";
import DiskBar from "@/components/echarts/diskBar.vue";
import { useRoute, useRouter } from 'vue-router'
import SocketService from '@/utils/websocket';


const cpu = ref("");
const gpu = ref("");
const memory = ref("");
const disk = ref("");


const route = useRoute()
const cpuTemChart = ref(null)
const gpuTemChart = ref(null)
const cpuUseChart = ref(null)
const gpuUseChart = ref(null)
const memoryUseChart = ref(null)
const diskUseChart = ref(null)

const state = {
  response: null,
  socket: null
}

const receiveData = (data) => {
  cpuTemChart.value.refreshData(data.cpuTemperature)
  gpuTemChart.value.refreshData(data.gpuTemperature)
  cpuUseChart.value.refreshData(data.time,data.cpuUse,data.cpuUseReal)
  gpuUseChart.value.refreshData(data.time,data.gpuUse,data.gpuUseReal)
  memoryUseChart.value.refreshData(data.time,data.memoryUse,data.memoryUseReal)
  diskUseChart.value.refreshData(data.diskPath,data.diskUse,data.diskFree)
}
// const url = "ws://localhost:8010/prom/GetHostRealInfo"
// const socket = new SocketService(url,receiveData)



// const chart = ref(null)

// setInterval(function () {
//   const random = +(Math.random() * 100).toFixed(2);
//   var a
//   a = chatData[0]
//   chatData.shift()
//   chatData.push(a)
//   console.log("父组件",chatData)
//   chart.value.refreshData(chatData)     
// }, 2000);

const init = () => {
  const instance = route.params.instance
  cpu.value = route.params.cpu
  gpu.value = route.params.gpu
  memory.value = route.params.memory
  disk.value = route.params.disk
  const serviceName = route.params.serviceName
  const url = "ws://localhost:8010/prom/GetHostRealInfo?instance="+instance+"&serviceName="+serviceName
  const socket = new SocketService(url,receiveData)
  socket.connect()
  // console.log(row)
}
init()

</script>


<style scoped>
.chart-container{
  position: relative;
  width: 100%;
  height: calc(100vh - 84px);
}

::v-deep .dashboard-line-box .dashboard-line{
  height: 100%;
  min-height: 360px;
}

</style>