import service from '@/utils/request'
import { ElMessage } from 'element-plus'

export const createOrder = (data) => {
  return service({
    url: '/order/createOrder',
    method: 'post',
    data
  })
}

export const deleteOrder = (data) => {
  return service({
    url: '/order/deleteOrder',
    method: 'delete',
    data
  })
}

export const deleteOrderByIds = (data) => {
  return service({
    url: '/order/deleteOrderByIds',
    method: 'delete',
    data
  })
}

export const updateOrder = (data) => {
  return service({
    url: '/order/updateOrder',
    method: 'put',
    data
  })
}

export const findOrder = (params) => {
  return service({
    url: '/order/findOrder',
    method: 'get',
    params
  })
}

export const getOrderList = (params) => {
  return service({
    url: '/order/getOrderList',
    method: 'get',
    params
  })
}

export const approveApply = (data) => {
  return service({
    url: '/order/approveApply',
    method: 'put',
    data
  })
}

export const rejectApply = (data) => {
  return service({
    url: '/order/rejectApply',
    method: 'put',
    data
  })
}

export const approveDelay = (data) => {
  return service({
    url: '/order/approveDelay',
    method: 'put',
    data
  })
}

export const rejectDelay = (data) => {
  return service({
    url: '/order/rejectDelay',
    method: 'put',
    data
  })
}

export const approveRefund = (data) => {
  return service({
    url: '/order/approveRefund',
    method: 'put',
    data
  })
}

export const rejectRefund = (data) => {
  return service({
    url: '/order/rejectRefund',
    method: 'put',
    data
  })
}

export const approveClaim = (data) => {
  return service({
    url: '/order/approveClaim',
    method: 'put',
    data
  })
}

export const rejectClaim = (data) => {
  return service({
    url: '/order/rejectClaim',
    method: 'put',
    data
  })
}

export const openLetter = (data) => {
  return service({
    url: '/order/openLetter',
    method: 'put',
    data
  })
}

export const rePush = (data) => {
  return service({
    url: '/order/rePush',
    method: 'put',
    data
  })
}

export const getOrderStatisticData = (params) => {
  return service({
    url: '/order/getOrderStatisticData',
    method: 'get',
    params
  })
}

export const getEmployeeStatisticData = (params) => {
  return service({
    url: '/order/getEmployeeStatisticData',
    method: 'get',
    params
  })
}

export const getOrderTrendData = (params) => {
  return service({
    url: '/order/getOrderTrendData',
    method: 'get',
    params
  })
}

export const downloadExcel = (params) => {
  return service({
    url: '/order/exportExcel',
    method: 'get',
    params: params,
    responseType: 'blob'
  }).then((res) => {
    handleExcelError(res)
  })
}

const handleExcelError = (res) => {
  if (typeof (res.data) !== 'undefined') {
    if (res.data.type === 'application/json') {
      const reader = new FileReader()
      reader.onload = function() {
        const message = JSON.parse(reader.result).msg
        ElMessage({
          showClose: true,
          message: message,
          type: 'error'
        })
      }
      reader.readAsText(new Blob([res.data]))
    }
  } else {
    const downloadUrl = window.URL.createObjectURL(new Blob([res]))
    const a = document.createElement('a')
    a.style.display = 'none'
    a.href = downloadUrl
    a.download = (new Date().getTime()).toString() + '.xlsx'
    a.click()
  }
}

export const findOrderByNos = (params) => {
  return service({
    url: '/order/findOrderByNos',
    method: 'get',
    params
  })
}

export const requestInvoice = (data) => {
  return service({
    url: '/order/requestInvoice',
    method: 'put',
    data
  })
}

export const queryInvoice = (data) => {
  return service({
    url: '/order/queryInvoice',
    method: 'put',
    data
  })
}

export const assignOrder = (data) => {
  return service({
    url: '/order/assignOrder',
    method: 'put',
    data
  })
}

export const markOfflineRefund = (data) => {
  return service({
    url: '/order/markOfflineRefund',
    method: 'post',
    data
  })
}

export const unmarkOfflineRefund = (data) => {
  return service({
    url: '/order/unmarkOfflineRefund',
    method: 'post',
    data
  })
}

export const downloadInvoiceExcel = (params) => {
  return service({
    url: '/order/exportInvoiceExcel',
    method: 'get',
    params: params,
    responseType: 'blob'
  }).then((res) => {
    handleExcelError(res)
  })
}

export const elogValidate = (data) => {
  return service({
    url: '/order/elogValidate',
    method: 'post',
    data
  })
}
