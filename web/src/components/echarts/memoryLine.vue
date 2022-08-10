<template>
  <!-- <div class="echarts-box">
    <div id="myEcharts" :style="{ width: this.width, height: this.height }"></div>
  </div> -->
    <div class="dashboard-line-box">
    <div class="dashboard-line-title">
      {{title}}   当前   {{real}} %
    </div>
    <div
      ref="echart"
      class="dashboard-line"
    />
  </div>
</template>

<script>
export default {
  props: {
    title: {
        type: String,          //类型
        default: 'kobe',       //设置默认值
        required: true         //是否必须传递
    }
    
  }
}
</script>


<script setup>
import * as echarts from 'echarts'
import { nextTick, onMounted, onUnmounted, ref , shallowRef , watch} from 'vue'

const real = ref(0)

const chart = shallowRef(null)
const echart = ref(null)
const initChart = () => {
  chart.value = echarts.init(echart.value, /*'macarons'*/)
  setOptions()
}
const setOptions = () => {
  chart.value.setOption({
    xAxis: {
    name: '时间',
    type: 'category',
    data: []
    },
    yAxis: {
      name: '百分比',
      type: 'value'
    },
  series: [
    {
      data: [],
      type: 'line'
    }
  ]
  })
}

onMounted(async() => {
  await nextTick()
  initChart()
})

onUnmounted(() => {
  if (!chart.value) {
    return
  }
  chart.value.dispose()
  chart.value = null
})

// const refreshData = ref(null)
const refreshData = (timeLine,data,r) =>{
  real.value = r
  chart.value.setOption({
    xAxis: [
      {
      data:timeLine
      }
    ],
    series: [
      {
      data:data
      }
    ]
  })
}


defineExpose({refreshData})

</script>




<style lang="scss" scoped>
.dashboard-line-box {
  .dashboard-line {
    background-color: #fff;
    height: 360px;
    width: 100%;
  }
  .dashboard-line-title {
    font-weight: 600;
    margin-bottom: 12px;
  }
}


</style>
