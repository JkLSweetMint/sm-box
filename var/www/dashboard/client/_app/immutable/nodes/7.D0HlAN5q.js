import{s,g as a,o as n}from"../chunks/scheduler.VUuLSjLj.js";import{S as i,i as c}from"../chunks/index.YcojCQiS.js";function l(t){const e=a("CallServiceMethod");return n(async()=>{let o=await e({service:"auth",url:"/basic-auth/logout",showToast:!1});o&&o.code==200?window.location.replace("/"):window.history.back()}),[]}class f extends i{constructor(e){super(),c(this,e,l,null,s,{})}}export{f as component};
