import{t as Ce,s as Y,d as Se,u as Pe,f as Ie,h as Te,c as W,x as Z,y as x,w as We,g as Ke,o as Je}from"./scheduler.B5rZiKyw.js";import{S as Ne,i as $e,s as j,k as N,f as V,g as y,u as I,z as De,v as T,y as Ae,d as g,e as O,c as C,a as S,F as A,G as q,h as G,E as le,l as w,n as ee,q as te,r as se,w as ne,t as Qe,b as Xe,j as Ye}from"./index.DrETqFkR.js";import{i as Ze,b as ae,t as ie,c as xe,d as et,w as tt,g as st,m as L,h as K,j as J,p as nt,k as rt,l as ue,n as z,q as M,r as ot,v as lt,x as Q,y as at,z as re,f as ce,e as fe,u as it,o as ut}from"./index.C2WeaqD0.js";import{s as ct,d as ft,e as F,u as dt,g as de,r as pt}from"./action.CNJ1jGTj.js";import{w as _t}from"./index.C_NBoP9Y.js";import{c as E}from"./index.Df_iRG8e.js";import{B as Le,N as Be,P as ht,C as gt,a as mt,G as vt}from"./Button.wSOItzZU.js";async function pe(a){const{prop:e,defaultEl:s}=a;if(await Promise.all([ct(1),Ce]),e===void 0){s==null||s.focus();return}const t=Ze(e)?e(s):e;if(typeof t=="string"){const n=document.querySelector(t);if(!ae(n))return;n.focus()}else ae(t)&&t.focus()}function bt(a,e){throw new Error("[MELTUI ERROR]: The `use:melt` action cannot be used without MeltUI's Preprocessor. See: https://www.melt-ui.com/docs/preprocessor")}const yt={positioning:{placement:"bottom"},arrowSize:8,defaultOpen:!1,disableFocusTrap:!1,escapeBehavior:"close",preventScroll:!1,onOpenChange:void 0,closeOnOutsideClick:!0,portal:"body",forceVisible:!1,openFocus:void 0,closeFocus:void 0,onOutsideClick:void 0,preventTextSelectionOverflow:!0},{name:B}=lt("popover"),wt=["trigger","content"];function kt(a){const e={...yt,...a},s=ie(xe(e,"open","ids")),{positioning:t,arrowSize:n,disableFocusTrap:l,preventScroll:o,escapeBehavior:i,closeOnOutsideClick:r,portal:u,forceVisible:c,openFocus:p,closeFocus:b,onOutsideClick:f,preventTextSelectionOverflow:v}=s,H=e.open??_t(e.defaultOpen),h=et(H,e==null?void 0:e.onOpenChange),$=tt.writable(null),D=ie({...st(wt),...e.ids});function oe(){h.set(!1)}const R=ft({open:h,activeTrigger:$,forceVisible:c}),Fe=L(B("content"),{stores:[R,h,$,u,D.content],returned:([d,m,_,k,P])=>({hidden:d&&K?void 0:!0,tabindex:-1,style:d?void 0:J({display:"none"}),id:P,"data-state":m&&_?"open":"closed","data-portal":nt(k)}),action:d=>{let m=Q;const _=F([R,$,t,l,r,u],([k,P,qe,Ge,He,Ue])=>{m(),!(!k||!P)&&Ce().then(()=>{m(),m=dt(d,{anchorElement:P,open:h,options:{floating:qe,focusTrap:Ge?null:void 0,modal:{shouldCloseOnInteractOutside:Re,onClose:oe,closeOnInteractOutside:He},escapeKeydown:{behaviorType:i},portal:de(d,Ue),preventTextSelectionOverflow:{enabled:v}}}).destroy})});return{destroy(){_(),m()}}}});async function U(){h.update(d=>!d)}function Re(d){var k;if((k=f.get())==null||k(d),d.defaultPrevented)return!1;const m=d.target,_=document.getElementById(D.trigger.get());return!(_&&rt(m)&&(m===_||_.contains(m)))}const ze=L(B("trigger"),{stores:[R,D.content,D.trigger],returned:([d,m,_])=>({role:"button","aria-haspopup":"dialog","aria-expanded":d?"true":"false","data-state":_e(d),"aria-controls":m,id:_}),action:d=>{$.set(d);const m=ue(z(d,"click",U),z(d,"keydown",_=>{_.key!==M.ENTER&&_.key!==M.SPACE||(_.preventDefault(),U())}));return{destroy(){$.set(null),m()}}}}),Me=L(B("overlay"),{stores:[R],returned:([d])=>({hidden:d?void 0:!0,tabindex:-1,style:J({display:d?void 0:"none"}),"aria-hidden":"true","data-state":_e(d)}),action:d=>{let m=Q,_=Q;return m=F([u],([k])=>{if(_(),k===null)return;const P=de(d,k);P!==null&&(_=ot(d,P).destroy)}),{destroy(){m(),_()}}}}),je=L(B("arrow"),{stores:n,returned:d=>({"data-arrow":!0,style:J({position:"absolute",width:`var(--arrow-size, ${d}px)`,height:`var(--arrow-size, ${d}px)`})})}),Ve=L(B("close"),{returned:()=>({type:"button"}),action:d=>({destroy:ue(z(d,"click",_=>{_.defaultPrevented||oe()}),z(d,"keydown",_=>{_.defaultPrevented||_.key!==M.ENTER&&_.key!==M.SPACE||(_.preventDefault(),U())}))})});return F([h,$,o],([d,m,_])=>{if(!K||!d)return;const k=[];return _&&k.push(pt()),pe({prop:p.get(),defaultEl:m}),()=>{k.forEach(P=>P())}}),F(h,d=>{if(!K||d)return;const m=document.getElementById(D.trigger.get());pe({prop:b.get(),defaultEl:m})},{skipFirstRun:!0}),{ids:D,elements:{trigger:ze,content:Fe,arrow:je,close:Ve,overlay:Me},states:{open:h},options:s}}function _e(a){return a?"open":"closed"}function he(a){return Object.keys(a)}function Et(a){let e={};return he(a).forEach(s=>{const t=a[s];F(t,n=>{var l;s in e&&((l=e[s])==null||l.call(e,n))})}),he(a).reduce((s,t)=>({...s,[t]:function(l,o){a[t].update(i=>at(i,l)?i:l),o&&(e={...e,[t]:o})}}),{})}const Ot=a=>({trigger:a&8}),ge=a=>({melt:bt,trigger:a[3]});function me(a){let e,s,t,n,l,o,i;const r=a[10].default,u=Se(r,a,a[9],null);let c=!(a[2]&X)&&ve(a),p=[a[4],{class:t="popover"+E(a[1])+E(a[2]&be?"no-border":"")+E(a[2]&ye?"no-shadow":"")}],b={};for(let f=0;f<p.length;f+=1)b=Z(b,p[f]);return{c(){e=O("div"),u&&u.c(),s=j(),c&&c.c(),this.h()},l(f){e=C(f,"DIV",{class:!0});var v=S(e);u&&u.l(v),s=V(v),c&&c.l(v),v.forEach(g),this.h()},h(){A(e,b),q(e,"svelte-4wje2y",!0)},m(f,v){y(f,e,v),u&&u.m(e,null),G(e,s),c&&c.m(e,null),l=!0,o||(i=x(a[4].action(e)),o=!0)},p(f,v){u&&u.p&&(!l||v&512)&&Pe(u,r,f,f[9],l?Te(r,f[9],v,null):Ie(f[9]),null),f[2]&X?c&&(c.d(1),c=null):c?c.p(f,v):(c=ve(f),c.c(),c.m(e,null)),A(e,b=re(p,[v&16&&f[4],(!l||v&6&&t!==(t="popover"+E(f[1])+E(f[2]&be?"no-border":"")+E(f[2]&ye?"no-shadow":"")))&&{class:t}])),q(e,"svelte-4wje2y",!0)},i(f){l||(I(u,f),f&&We(()=>{l&&(n||(n=le(e,ce,{duration:300,y:100},!0)),n.run(1))}),l=!0)},o(f){T(u,f),f&&(n||(n=le(e,ce,{duration:300,y:100},!1)),n.run(0)),l=!1},d(f){f&&g(e),u&&u.d(f),c&&c.d(),f&&n&&n.end(),o=!1,i()}}}function ve(a){let e,s,t,n,l=[{type:"button"},{class:"close"},a[5]],o={};for(let i=0;i<l.length;i+=1)o=Z(o,l[i]);return{c(){e=O("button"),s=O("i"),this.h()},l(i){e=C(i,"BUTTON",{type:!0,class:!0});var r=S(e);s=C(r,"I",{class:!0}),S(s).forEach(g),r.forEach(g),this.h()},h(){w(s,"class","fa-solid fa-xmark"),A(e,o),q(e,"svelte-4wje2y",!0)},m(i,r){y(i,e,r),G(e,s),e.autofocus&&e.focus(),t||(n=x(a[5].action(e)),t=!0)},p(i,r){A(e,o=re(l,[{type:"button"},{class:"close"},r&32&&i[5]])),q(e,"svelte-4wje2y",!0)},d(i){i&&g(e),t=!1,n()}}}function Ct(a){let e,s,t;const n=a[10].trigger,l=Se(n,a,a[9],ge);let o=a[0]&&me(a);return{c(){l&&l.c(),e=j(),o&&o.c(),s=N()},l(i){l&&l.l(i),e=V(i),o&&o.l(i),s=N()},m(i,r){l&&l.m(i,r),y(i,e,r),o&&o.m(i,r),y(i,s,r),t=!0},p(i,[r]){l&&l.p&&(!t||r&520)&&Pe(l,n,i,i[9],t?Te(n,i[9],r,Ot):Ie(i[9]),ge),i[0]?o?(o.p(i,r),r&1&&I(o,1)):(o=me(i),o.c(),I(o,1),o.m(s.parentNode,s)):o&&(De(),T(o,1,1,()=>{o=null}),Ae())},i(i){t||(I(l,i),I(o),t=!0)},o(i){T(l,i),T(o),t=!1},d(i){i&&(g(e),g(s)),l&&l.d(i),o&&o.d(i)}}}const X=1,be=2,ye=4;function St(a,e,s){let t,n,l,{$$slots:o={},$$scope:i}=e,{class:r=""}=e,{open:u=!1}=e,{flags:c=0}=e;const{elements:{trigger:p,content:b,close:f},states:v}=kt({forceVisible:!0});W(a,p,h=>s(3,t=h)),W(a,b,h=>s(4,n=h)),W(a,f,h=>s(5,l=h));const H=Et(v);return a.$$set=h=>{"class"in h&&s(1,r=h.class),"open"in h&&s(0,u=h.open),"flags"in h&&s(2,c=h.flags),"$$scope"in h&&s(9,i=h.$$scope)},a.$$.update=()=>{a.$$.dirty&1&&H.open(u,h=>s(0,u=h))},[u,r,c,t,n,l,p,b,f,i,o]}class Pt extends Ne{constructor(e){super(),$e(this,e,St,Ct,Y,{class:1,open:0,flags:2})}}function we(a,e,s){const t=a.slice();t[11]=e[s],t[13]=s;const n=t[5](t[11].code);return t[10]=n,t}function It(a){const e=a.slice(),s=e[5](e[3]);return e[10]=s,e}function Tt(a){let e,s,t,n,l=a[11].name+"",o,i;return{c(){e=O("span"),t=j(),n=O("span"),o=Qe(l),i=j(),this.h()},l(r){e=C(r,"SPAN",{class:!0}),S(e).forEach(g),t=V(r),n=C(r,"SPAN",{class:!0});var u=S(n);o=Xe(u,l),u.forEach(g),i=V(r),this.h()},h(){w(e,"class",s=`fi fi-${a[10]} fis rounded !w-5 !h-5 saturate-[1.25]`+E(a[11].active?"":"saturate-[.25]")),w(n,"class","mr-auto max-w-24 break-words")},m(r,u){y(r,e,u),y(r,t,u),y(r,n,u),G(n,o),y(r,i,u)},p(r,u){u&4&&s!==(s=`fi fi-${r[10]} fis rounded !w-5 !h-5 saturate-[1.25]`+E(r[11].active?"":"saturate-[.25]"))&&w(e,"class",s),u&4&&l!==(l=r[11].name+"")&&Ye(o,l)},d(r){r&&(g(e),g(t),g(n),g(i))}}}function ke(a,e){let s,t,n;function l(){return e[6](e[11])}return t=new Le({props:{class:"hover:!bg-black/5 focus-visible:!bg-black/5 !rounded-none",flags:mt|Be,palette:vt,disabled:!e[11].active,OnClick:l,$$slots:{default:[Tt]},$$scope:{ctx:e}}}),{key:a,first:null,c(){s=N(),ee(t.$$.fragment),this.h()},l(o){s=N(),te(t.$$.fragment,o),this.h()},h(){this.first=s},m(o,i){y(o,s,i),se(t,o,i),n=!0},p(o,i){e=o;const r={};i&4&&(r.disabled=!e[11].active),i&4&&(r.OnClick=l),i&16388&&(r.$$scope={dirty:i,ctx:e}),t.$set(r)},i(o){n||(I(t.$$.fragment,o),n=!0)},o(o){T(t.$$.fragment,o),n=!1},d(o){o&&g(s),ne(t,o)}}}function Nt(a){let e,s,t=[],n=new Map,l,o=fe(a[2]);const i=r=>r[13];for(let r=0;r<o.length;r+=1){let u=we(a,o,r),c=i(u);n.set(c,t[r]=ke(c,u))}return{c(){e=O("div"),s=O("div");for(let r=0;r<t.length;r+=1)t[r].c();this.h()},l(r){e=C(r,"DIV",{class:!0});var u=S(e);s=C(u,"DIV",{class:!0});var c=S(s);for(let p=0;p<t.length;p+=1)t[p].l(c);c.forEach(g),u.forEach(g),this.h()},h(){w(s,"class","flex flex-col w-40"),w(e,"class","max-h-80 overflow-y-auto")},m(r,u){y(r,e,u),G(e,s);for(let c=0;c<t.length;c+=1)t[c]&&t[c].m(s,null);l=!0},p(r,u){u&52&&(o=fe(r[2]),De(),t=it(t,u,i,1,r,o,n,s,ut,ke,null,we),Ae())},i(r){if(!l){for(let u=0;u<o.length;u+=1)I(t[u]);l=!0}},o(r){for(let u=0;u<t.length;u+=1)T(t[u]);l=!1},d(r){r&&g(e);for(let u=0;u<t.length;u+=1)t[u].d()}}}function $t(a){let e,s;return{c(){e=O("div"),this.h()},l(t){e=C(t,"DIV",{class:!0,style:!0}),S(e).forEach(g),this.h()},h(){w(e,"class","rounded bg-black/20 animate-pulse duration-500 m-1"),w(e,"style",s=`width: ${a[1]}; height: ${a[1]};`)},m(t,n){y(t,e,n)},p(t,n){n&2&&s!==(s=`width: ${t[1]}; height: ${t[1]};`)&&w(e,"style",s)},d(t){t&&g(e)}}}function Ee(a){let e,s,t;return{c(){e=O("span"),this.h()},l(n){e=C(n,"SPAN",{class:!0,style:!0}),S(e).forEach(g),this.h()},h(){w(e,"class",s=`fi fi-${a[10]} fis rounded saturate-[1.25] m-1`),w(e,"style",t=`width: ${a[1]}; height: ${a[1]};`)},m(n,l){y(n,e,l)},p(n,l){l&8&&s!==(s=`fi fi-${n[10]} fis rounded saturate-[1.25] m-1`)&&w(e,"class",s),l&2&&t!==(t=`width: ${n[1]}; height: ${n[1]};`)&&w(e,"style",t)},d(n){n&&g(e)}}}function Oe(a){let e;function s(o,i){return o[3]?Ee:$t}function t(o,i){return i===Ee?It(o):o}let n=s(a),l=n(t(a,n));return{c(){l.c(),e=N()},l(o){l.l(o),e=N()},m(o,i){l.m(o,i),y(o,e,i)},p(o,i){n===(n=s(o))&&l?l.p(t(o,n),i):(l.d(1),l=n(t(o,n)),l&&(l.c(),l.m(e.parentNode,e)))},d(o){o&&g(e),l.d(o)}}}function Dt(a){let e=a[3],s,t=Oe(a);return{c(){t.c(),s=N()},l(n){t.l(n),s=N()},m(n,l){t.m(n,l),y(n,s,l)},p(n,l){l&8&&Y(e,e=n[3])?(t.d(1),t=Oe(n),t.c(),t.m(s.parentNode,s)):t.p(n,l)},d(n){n&&g(s),t.d(n)}}}function At(a){let e,s,t,n,l;s=new Le({props:{class:"!p-1 w-max h-max"+E(a[0]),flags:Be|ht|gt,$$slots:{default:[Dt]},$$scope:{ctx:a}}});let o=[{slot:"trigger"},a[9]],i={};for(let r=0;r<o.length;r+=1)i=Z(i,o[r]);return{c(){e=O("div"),ee(s.$$.fragment),this.h()},l(r){e=C(r,"DIV",{slot:!0});var u=S(e);te(s.$$.fragment,u),u.forEach(g),this.h()},h(){A(e,i)},m(r,u){y(r,e,u),se(s,e,null),t=!0,n||(l=x(a[9].action(e)),n=!0)},p(r,u){const c={};u&1&&(c.class="!p-1 w-max h-max"+E(r[0])),u&16394&&(c.$$scope={dirty:u,ctx:r}),s.$set(c),A(e,i=re(o,[{slot:"trigger"},u&512&&r[9]]))},i(r){t||(I(s.$$.fragment,r),t=!0)},o(r){T(s.$$.fragment,r),t=!1},d(r){r&&g(e),ne(s),n=!1,l()}}}function Lt(a){let e,s;return e=new Pt({props:{class:"!p-0 overflow-hidden",flags:X,$$slots:{trigger:[At,({melt:t,trigger:n})=>({8:t,9:n}),({melt:t,trigger:n})=>(t?256:0)|(n?512:0)],default:[Nt]},$$scope:{ctx:a}}}),{c(){ee(e.$$.fragment)},l(t){te(e.$$.fragment,t)},m(t,n){se(e,t,n),s=!0},p(t,[n]){const l={};n&16911&&(l.$$scope={dirty:n,ctx:t}),e.$set(l)},i(t){s||(I(e.$$.fragment,t),s=!0)},o(t){T(e.$$.fragment,t),s=!1},d(t){ne(e,t)}}}function Bt(a,e,s){let t,n="";const l=Ke("CallServiceMethod"),o=async function(p){if(p==n)return;let b=await l({service:"i18n",url:"/languages/set",headers:{"Content-Type":"application/json;charset=utf-8"},method:"POST",data:{code:p},showToast:!0});b&&b.code==200&&window.location.reload()};let{class:i=""}=e,{size:r="24px"}=e;const u=function(p){return(p.split("-").slice(-1)[0]??"").toLowerCase()};Je(async()=>{var b,f;let p=await l({service:"i18n",url:"/languages/select"});p&&p.code==200&&(s(2,t=((b=p.data)==null?void 0:b.list)??[]),s(3,n=((f=p.data)==null?void 0:f.current)??""))});const c=p=>o(p.code);return a.$$set=p=>{"class"in p&&s(0,i=p.class),"size"in p&&s(1,r=p.size)},[i,r,t,n,o,u,c]}class Gt extends Ne{constructor(e){super(),$e(this,e,Bt,Lt,Y,{class:0,size:1})}}export{X as D,Gt as L,Pt as P,pe as h};