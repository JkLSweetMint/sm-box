import{s as E,d as L,u as S,f as q,h as G,c as A,k as h,o as M,n as D,a as F}from"../chunks/scheduler.B4NQwY3t.js";import{S as I,i as N,n as O,j as k,q as H,k as W,r as j,a as U,u as w,v,d as x,w as z}from"../chunks/index.gMR2UOS6.js";import{r as p}from"../chunks/index.DtvkMZOO.js";import{R as B,G as g,F as J,W as K,s as R,E as T,C}from"../chunks/index.BX9MmS-n.js";import{G as P}from"../chunks/Text.Dzl9t-jU.js";function Q(n){let a,o,r;a=new B({});const c=n[3].default,t=L(c,n,n[2],null);return{c(){O(a.$$.fragment),o=k(),t&&t.c()},l(e){H(a.$$.fragment,e),o=W(e),t&&t.l(e)},m(e,i){j(a,e,i),U(e,o,i),t&&t.m(e,i),r=!0},p(e,[i]){t&&t.p&&(!r||i&4)&&S(t,c,e,e[2],r?G(c,e[2],i,null):q(e[2]),null)},i(e){r||(w(a.$$.fragment,e),w(t,e),r=!0)},o(e){v(a.$$.fragment,e),v(t,e),r=!1},d(e){e&&x(o),z(a,e),t&&t.d(e)}}}function V(n,a,o){let r,c,t=D,e=()=>(t(),t=F(m,s=>o(5,c=s)),m);n.$$.on_destroy.push(()=>t());let{$$slots:i={},$$scope:y}=a,m=p();e();const _=p({auth:g({name:"authentication"}),i18n:g({name:"i18n"})});A(n,_,s=>o(4,r=s)),h("services",_);const $=async function(s){var l;const u=P("dashboard.toasts.error.title",c)||"An error occured";try{const d=r[s.service];if(!d)throw new Error(`Failed to call service method: service "${d}" is not defined`);let f=await J({url:`${d}${s.url}`,method:s.method,query:s.query,headers:s.headers,data:s.data,flags:K});return f.data.code!=200&&s.showToast&&R({data:{title:u,description:`${(l=f.data.error)==null?void 0:l.message}`,type:T,flags:C}}),f.data.headers=f.meta.headers,f.data}catch(d){return s.showToast&&R({data:{title:u,description:`${d}`,type:T,flags:C}}),{code:500,code_message:"",status:"fatal"}}};h("CallServiceMethod",$);const b=async function(s){var l;let u=await $({service:"i18n",url:"/texts/dictionary",query:{paths:s},showToast:!0});if(!(!u||u.code!=200))return p(((l=u.data)==null?void 0:l.dictionary)??{})};return h("i18nLoadDictionary",b),M(async()=>{e(o(0,m=await b("dashboard.toasts")??p()))}),n.$$set=s=>{"$$scope"in s&&o(2,y=s.$$scope)},[m,_,y,i]}class se extends I{constructor(a){super(),N(this,a,V,Q,E,{})}}export{se as component};