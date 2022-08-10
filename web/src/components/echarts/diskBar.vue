<template>
  <!-- <div class="echarts-box">
    <div id="myEcharts" :style="{ width: this.width, height: this.height }"></div>
  </div> -->
    <div class="dashboard-line-box">
    <div class="dashboard-line-title">
      {{title}}
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



const chart = shallowRef(null)
const echart = ref(null)
const initChart = () => {
  chart.value = echarts.init(echart.value, /*'macarons'*/)
  setOptions()
}
const setOptions = () => {
  chart.value.setOption({
//
    tooltip: {
    trigger: 'axis',
    axisPointer: {
      // Use axis to trigger tooltip
      type: 'shadow' // 'shadow' as default; can also be 'line' or 'shadow'
    }
  },
  legend: {},
  grid: {
    left: '3%',
    right: '4%',
    bottom: '3%',
    containLabel: true
  },
  xAxis: {
    name: '容量(T)',
    type: 'value'
  },
  yAxis: {
    name: '磁盘',
    type: 'category',
    data: []
  },
  series: [
    {
      name: 'Used',
      //color: '#CC0033',
      type: 'bar',
      stack: 'total',
      label: {
        show: true
      },
      emphasis: {
        focus: 'series'
      },
      data: []
    },
    {
      name: 'Free',
      //color: '#99CC00',
      type: 'bar',
      stack: 'total',
      label: {
        show: true
      },
      emphasis: {
        focus: 'series'
      },
      data: []
    }
  ]
//

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
const refreshData = (paths,used,free) =>{
  chart.value.setOption({
    yAxis:[
      {
        data:paths
      }
    ],
    series:[
      {
        data:used
      },
      {
        data:free
      }
    ]
  })
}


defineExpose({refreshData})


</script>


<style lang="scss" scoped>
.dashboard-line-box {
  .dashboard-line {
    background-color: rgb(250, 246, 246);
    height: 360px;
    width: 100%;
  }
  .dashboard-line-title {
    font-weight: 600;
    margin-bottom: 12px;
  }
}


</style>
