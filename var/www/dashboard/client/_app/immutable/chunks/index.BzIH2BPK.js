import{n as O,m as Ft,i as Gt,s as At,C as K,D as W,l as Ht,r as Mt,c as ut,o as Pt,a as rt}from"./scheduler.VUuLSjLj.js";import{B as jt,C as Wt,D as zt,E as Jt,S as xt,i as Ct,b as x,j as Y,H as Rt,t as St,f as C,g as R,d as k,k as Z,F as Lt,c as Nt,h as B,G as V,I as U,a as tt,l as I,J as Xt,K as Kt,L as Qt,s as Yt,e as vt,z as Zt,x as $t,u as Vt,v as Ut,n as te,q as ee,r as se,w as ne}from"./index.YcojCQiS.js";import{c as ae,t as Bt,o as oe,a as re,m as $,b as Ot,u as le,e as bt,d as lt,g as ft,n as ie,i as wt,k as Et,f as it,h as Q,j as kt,l as Dt,p as ce,q as ue}from"./index.DbE-JpqF.js";import{w as _t,d as de,a as fe}from"./index.oxanXEPY.js";function _e(s,e,t,n){if(!e)return O;const a=s.getBoundingClientRect();if(e.left===a.left&&e.right===a.right&&e.top===a.top&&e.bottom===a.bottom)return O;const{delay:d=0,duration:_=300,easing:c=Ft,start:i=jt()+d,end:r=i+_,tick:f=O,css:p}=t(s,{from:e,to:a},n);let w=!0,v=!1,l;function m(){p&&(l=zt(s,0,1,_,d,c,p)),d||(v=!0)}function h(){p&&Jt(s,l),w=!1}return Wt(u=>{if(!v&&u>=i&&(v=!0),v&&u>=r&&(f(1,0),h()),!w)return!1;if(v){const g=u-i,b=0+1*c(g/_);f(b,1-b)}return!0}),m(),f(0,1),h}function pe(s){const e=getComputedStyle(s);if(e.position!=="absolute"&&e.position!=="fixed"){const{width:t,height:n}=e,a=s.getBoundingClientRect();s.style.position="absolute",s.style.width=t,s.style.height=n,me(s,a)}}function me(s,e){const t=s.getBoundingClientRect();if(e.left!==t.left||e.top!==t.top){const n=getComputedStyle(s),a=n.transform==="none"?"":n.transform;s.style.transform=`${a} translate(${e.left-t.left}px, ${e.top-t.top}px)`}}function he(s,{from:e,to:t},n={}){const a=getComputedStyle(s),d=a.transform==="none"?"":a.transform,[_,c]=a.transformOrigin.split(" ").map(parseFloat),i=e.left+e.width*_/t.width-(t.left+_),r=e.top+e.height*c/t.height-(t.top+c),{delay:f=0,duration:p=v=>Math.sqrt(v)*120,easing:w=ae}=n;return{delay:f,duration:Gt(p)?p(Math.sqrt(i*i+r*r)):p,easing:w,css:(v,l)=>{const m=l*i,h=l*r,u=v+l*e.width/t.width,g=v+l*e.height/t.height;return`transform: ${d} translate(${m}px, ${h}px) scale(${u}, ${g});`}}}const ge={defaultValue:0,max:100},{name:ye}=Ot("progress"),ve=s=>{const e={...ge,...s},t=Bt(oe(e,"value")),{max:n}=t,a=e.value??_t(e.defaultValue),d=re(a,e==null?void 0:e.onValueChange);return{elements:{root:$(ye(),{stores:[d,n],returned:([c,i])=>({value:c,max:i,role:"meter","aria-valuemin":0,"aria-valuemax":i,"aria-valuenow":c,"data-value":c,"data-state":c===null?"indeterminate":c===i?"complete":"loading","data-max":i})})},states:{value:d},options:t}},{name:ct}=Ot("toast"),be={closeDelay:5e3,type:"foreground"};function we(s){const e={...be,...s},t=Bt(e),{closeDelay:n,type:a}=t,d=_t(new Map),_=l=>{const m={closeDelay:n.get(),type:a.get(),...l},h={content:ft(),title:ft(),description:ft()},u=m.closeDelay===0?null:window.setTimeout(()=>{c(h.content)},m.closeDelay),g=()=>{const{createdAt:D,pauseDuration:A,closeDelay:L,pausedAt:S}=b;return L===0?0:S?100*(S-D-A)/L:100*(performance.now()-D-A)/L},b={id:h.content,ids:h,...m,timeout:u,createdAt:performance.now(),pauseDuration:0,getPercentage:g};return d.update(D=>(D.set(h.content,b),new Map(D))),b},c=l=>{d.update(m=>(m.delete(l),new Map(m)))},i=(l,m)=>{d.update(h=>{const u=h.get(l);return u?(h.set(l,{...u,data:m}),new Map(h)):h})},r=$(ct("content"),{stores:d,returned:l=>m=>{const h=l.get(m);if(!h)return null;const{...u}=h;return{id:m,role:"alert","aria-describedby":u.ids.description,"aria-labelledby":u.ids.title,"aria-live":u.type==="foreground"?"assertive":"polite",tabindex:-1}},action:l=>{let m=ie;return m=bt(lt(l,"pointerenter",h=>{wt(h)||d.update(u=>{const g=u.get(l.id);return!g||g.closeDelay===0?u:(g.timeout!==null&&window.clearTimeout(g.timeout),g.pausedAt=performance.now(),new Map(u))})}),lt(l,"pointerleave",h=>{wt(h)||d.update(u=>{const g=u.get(l.id);if(!g||g.closeDelay===0)return u;const b=g.pausedAt??g.createdAt,D=b-g.createdAt-g.pauseDuration,A=g.closeDelay-D;return g.timeout=window.setTimeout(()=>{c(l.id)},A),g.pauseDuration+=performance.now()-b,g.pausedAt=void 0,new Map(u)})}),()=>{c(l.id)}),{destroy:m}}}),f=$(ct("title"),{stores:d,returned:l=>m=>{const h=l.get(m);return h?{id:h.ids.title}:null}}),p=$(ct("description"),{stores:d,returned:l=>m=>{const h=l.get(m);return h?{id:h.ids.description}:null}}),w=$(ct("close"),{returned:()=>l=>({type:"button","data-id":l}),action:l=>{function m(){l.dataset.id&&c(l.dataset.id)}return{destroy:bt(lt(l,"click",()=>{m()}),lt(l,"keydown",u=>{u.key!==Et.ENTER&&u.key!==Et.SPACE||(u.preventDefault(),m())}))}}}),v=de(d,l=>Array.from(l.values()));return{elements:{content:r,title:f,description:p,close:w},states:{toasts:fe(v)},helpers:{addToast:_,removeToast:c,updateToast:i},actions:{portal:le},options:t}}function Ee(s){let e=s[4].description+"",t;return{c(){t=St(e)},l(n){t=Nt(n,e)},m(n,a){tt(n,t,a)},p(n,a){a&16&&e!==(e=n[4].description+"")&&Yt(t,e)},d(n){n&&k(t)}}}function ke(s){let e,t=s[4].description+"",n;return{c(){e=new Rt(!1),n=vt(),this.h()},l(a){e=Lt(a,!1),n=vt(),this.h()},h(){e.a=n},m(a,d){e.m(t,a,d),tt(a,n,d)},p(a,d){d&16&&t!==(t=a[4].description+"")&&e.p(t)},d(a){a&&(k(n),e.d())}}}function De(s){let e,t,n,a,d,_,c,i,r,f=s[15](s[4].type)+"",p,w,v,l=s[4].title+"",m,h,u,g,b,D,A,L,S,T,z,J,P=[s[9],{class:"toast__indicator"}],F={};for(let o=0;o<P.length;o+=1)F=K(F,P[o]);let G=[s[2],{class:"flex items-center gap-2 text-lg font-semibold"}],N={};for(let o=0;o<G.length;o+=1)N=K(N,G[o]);function H(o,E){return o[4].flags&Ie?ke:Ee}let M=H(s),q=M(s),X=[s[1],{class:"text-sm"}],y={};for(let o=0;o<X.length;o+=1)y=K(y,X[o]);let j=[s[0],{class:"toast__close"}],et={};for(let o=0;o<j.length;o+=1)et=K(et,j[o]);let dt=[s[3],{class:A="toast"+it(s[4].type)+it(s[4].flags&It?"with-main-color":"")}],st={};for(let o=0;o<dt.length;o+=1)st=K(st,dt[o]);return{c(){e=x("div"),t=x("div"),n=x("div"),d=Y(),_=x("div"),c=x("div"),i=x("div"),r=new Rt(!1),p=Y(),w=x("div"),v=x("span"),m=St(l),h=Y(),u=x("span"),q.c(),g=Y(),b=x("button"),D=x("i"),this.h()},l(o){e=C(o,"DIV",{class:!0});var E=R(e);t=C(E,"DIV",{class:!0});var pt=R(t);n=C(pt,"DIV",{style:!0,class:!0}),R(n).forEach(k),pt.forEach(k),d=Z(E),_=C(E,"DIV",{class:!0});var nt=R(_);c=C(nt,"DIV",{class:!0});var at=R(c);i=C(at,"DIV",{class:!0});var mt=R(i);r=Lt(mt,!1),mt.forEach(k),p=Z(at),w=C(at,"DIV",{class:!0});var ot=R(w);v=C(ot,"SPAN",{class:!0});var ht=R(v);m=Nt(ht,l),ht.forEach(k),h=Z(ot),u=C(ot,"SPAN",{class:!0});var gt=R(u);q.l(gt),gt.forEach(k),ot.forEach(k),at.forEach(k),g=Z(nt),b=C(nt,"BUTTON",{class:!0});var yt=R(b);D=C(yt,"I",{class:!0}),R(D).forEach(k),yt.forEach(k),nt.forEach(k),E.forEach(k),this.h()},h(){B(n,"style",a=`transform: translateX(-${100*(s[10]??0)/(s[11]??1)}%)`),B(n,"class","svelte-14k9omq"),V(t,F),U(t,"svelte-14k9omq",!0),r.a=null,B(i,"class","flex flex-col items-center justify-center mr-3"),V(v,N),U(v,"svelte-14k9omq",!0),V(u,y),U(u,"svelte-14k9omq",!0),B(w,"class","flex flex-col"),B(c,"class","flex flex-row"),B(D,"class","fa-solid fa-xmark"),V(b,et),U(b,"svelte-14k9omq",!0),B(_,"class","toast__content svelte-14k9omq"),V(e,st),U(e,"svelte-14k9omq",!0)},m(o,E){tt(o,e,E),I(e,t),I(t,n),I(e,d),I(e,_),I(_,c),I(c,i),r.m(f,i),I(c,p),I(c,w),I(w,v),I(v,m),I(w,h),I(w,u),q.m(u,null),I(_,g),I(_,b),I(b,D),b.autofocus&&b.focus(),T=!0,z||(J=[W(s[9].action(t)),W(s[2].action(v)),W(s[1].action(u)),W(s[0].action(b)),W(s[3].action(e))],z=!0)},p(o,[E]){(!T||E&3072&&a!==(a=`transform: translateX(-${100*(o[10]??0)/(o[11]??1)}%)`))&&B(n,"style",a),V(t,F=Q(P,[E&512&&o[9],{class:"toast__indicator"}])),U(t,"svelte-14k9omq",!0),(!T||E&16)&&f!==(f=o[15](o[4].type)+"")&&r.p(f),(!T||E&16)&&l!==(l=o[4].title+"")&&Xt(m,l,N.contenteditable),V(v,N=Q(G,[E&4&&o[2],{class:"flex items-center gap-2 text-lg font-semibold"}])),U(v,"svelte-14k9omq",!0),M===(M=H(o))&&q?q.p(o,E):(q.d(1),q=M(o),q&&(q.c(),q.m(u,null))),V(u,y=Q(X,[E&2&&o[1],{class:"text-sm"}])),U(u,"svelte-14k9omq",!0),V(b,et=Q(j,[E&1&&o[0],{class:"toast__close"}])),U(b,"svelte-14k9omq",!0),V(e,st=Q(dt,[E&8&&o[3],(!T||E&16&&A!==(A="toast"+it(o[4].type)+it(o[4].flags&It?"with-main-color":"")))&&{class:A}])),U(e,"svelte-14k9omq",!0)},i(o){T||(o&&Ht(()=>{T&&(S&&S.end(1),L=Kt(e,kt,{duration:150,x:"100%"}),L.start())}),T=!0)},o(o){L&&L.invalidate(),o&&(S=Qt(e,kt,{duration:150,x:"100%"})),T=!1},d(o){o&&k(e),q.d(),o&&S&&S.end(),z=!1,Mt(J)}}}const It=1,Ie=2,Te="info",qe="success",Ae="warning",xe="error";function Ce(s,e,t){let n,a,d,_,c,i,r,f,p,w,v,l,m=O,h=()=>(m(),m=rt(_,y=>t(19,l=y)),_),u,g=O,b=()=>(g(),g=rt(d,y=>t(20,u=y)),d),D,A=O,L=()=>(A(),A=rt(a,y=>t(21,D=y)),a),S,T=O,z=()=>(T(),T=rt(n,y=>t(22,S=y)),n),J,P,F;s.$$.on_destroy.push(()=>m()),s.$$.on_destroy.push(()=>g()),s.$$.on_destroy.push(()=>A()),s.$$.on_destroy.push(()=>T());let{elements:G}=e,{toast:N}=e;const H=_t(0);ut(s,H,y=>t(10,P=y));const{elements:{root:M},options:{max:q}}=ve({max:100,value:H});ut(s,M,y=>t(9,J=y)),ut(s,q,y=>t(11,F=y));const X=function(y){switch(y){case Te:return'<i class="fa-solid fa-circle-info text-info text-xl"></i>';case qe:return'<i class="fa-solid fa-circle-check text-success text-xl"></i>';case Ae:return'<i class="fa-solid fa-triangle-exclamation text-warning text-xl"></i>';case xe:return'<i class="fa-solid fa-circle-exclamation text-error text-xl"></i>'}return""};return Pt(()=>{let y;const j=()=>{H.set(r()),y=requestAnimationFrame(j)};return y=requestAnimationFrame(j),()=>cancelAnimationFrame(y)}),s.$$set=y=>{"elements"in y&&t(16,G=y.elements),"toast"in y&&t(17,N=y.toast)},s.$$.update=()=>{s.$$.dirty&65536&&z(t(8,{content:n,title:a,description:d,close:_}=G,n,L(t(7,a)),b(t(6,d)),h(t(5,_)))),s.$$.dirty&131072&&t(4,{data:c,id:i,getPercentage:r}=N,c,(t(18,i),t(17,N))),s.$$.dirty&4456448&&t(3,f=S(i)),s.$$.dirty&2359296&&t(2,p=D(i)),s.$$.dirty&1310720&&t(1,w=u(i)),s.$$.dirty&786432&&t(0,v=l(i))},[v,w,p,f,c,_,d,a,n,J,P,F,H,M,q,X,G,N,i,l,u,D,S]}class Re extends xt{constructor(e){super(),Ct(this,e,Ce,De,At,{elements:16,toast:17})}}function Tt(s,e,t){const n=s.slice();return n[1]=e[t],n}function qt(s,e){let t,n,a,d,_=O,c;return n=new Re({props:{elements:Le,toast:e[1]}}),{key:s,first:null,c(){t=x("div"),te(n.$$.fragment),a=Y(),this.h()},l(i){t=C(i,"DIV",{});var r=R(t);ee(n.$$.fragment,r),a=Z(r),r.forEach(k),this.h()},h(){this.first=t},m(i,r){tt(i,t,r),se(n,t,null),I(t,a),c=!0},p(i,r){e=i;const f={};r&1&&(f.toast=e[1]),n.$set(f)},r(){d=t.getBoundingClientRect()},f(){pe(t),_()},a(){_(),_=_e(t,d,he,{duration:500})},i(i){c||(Vt(n.$$.fragment,i),c=!0)},o(i){Ut(n.$$.fragment,i),c=!1},d(i){i&&k(t),ne(n)}}}function Se(s){let e,t=[],n=new Map,a,d,_,c=Dt(s[0]);const i=r=>r[1].id;for(let r=0;r<c.length;r+=1){let f=Tt(s,c,r),p=i(f);n.set(p,t[r]=qt(p,f))}return{c(){e=x("div");for(let r=0;r<t.length;r+=1)t[r].c();this.h()},l(r){e=C(r,"DIV",{class:!0});var f=R(e);for(let p=0;p<t.length;p+=1)t[p].l(f);f.forEach(k),this.h()},h(){B(e,"class","toast-container svelte-zln6l1")},m(r,f){tt(r,e,f);for(let p=0;p<t.length;p+=1)t[p]&&t[p].m(e,null);a=!0,d||(_=W(Ue.call(null,e)),d=!0)},p(r,[f]){if(f&1){c=Dt(r[0]),Zt();for(let p=0;p<t.length;p+=1)t[p].r();t=ce(t,f,i,1,r,c,n,e,ue,qt,null,Tt);for(let p=0;p<t.length;p+=1)t[p].a();$t()}},i(r){if(!a){for(let f=0;f<c.length;f+=1)Vt(t[f]);a=!0}},o(r){for(let f=0;f<t.length;f+=1)Ut(t[f]);a=!1},d(r){r&&k(e);for(let f=0;f<t.length;f+=1)t[f].d();d=!1,_()}}}const{elements:Le,helpers:{addToast:Ne},states:{toasts:Ve},actions:{portal:Ue}}=we(),Pe=Ne;function Be(s,e,t){let n;return ut(s,Ve,a=>t(0,n=a)),[n]}class je extends xt{constructor(e){super(),Ct(this,e,Be,Se,At,{})}}const Oe=1,We=s=>`/${s.name}/api/v${s.version??"1.0"}`,ze=async function(s){s.flags=s.flags?s.flags:0;try{let e={method:s.method??"GET",credentials:s.flags&Oe?"include":"omit",headers:s.headers};e.method!="GET"&&(e.body=JSON.stringify(s.data));let t="";s.query&&(t="?"+Object.entries(s.query).map(_=>Array.isArray(_[1])?_[1].map(c=>`${_[0]}=${c}`):[`${_[0]}=${_[1]}`]).map(_=>_.join("&")).join("&"));const n=await fetch(s.url+t,e),a=Object.fromEntries(n.headers.entries());return{data:await n.json(),meta:{headers:a}}}catch(e){throw new Error(`Failed to call a request: ${e}`)}},Je=async function(s){return(await fetch(window.location.toString(),{method:"HEAD"})).headers.get(s)||null};export{It as C,xe as E,ze as F,We as G,je as R,Oe as W,Je as a,Pe as s};