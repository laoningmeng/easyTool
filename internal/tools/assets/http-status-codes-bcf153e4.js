import{d as k,z as B,an as b,bF as i,e as F,j as n,k as w,n as p,B as C,F as m,r as h,C as S,o as t,q as c,s,E as z,w as E,y as N}from"./index-a828413e.js";const R={"mb-2":"","text-xl":""},T={"text-lg":"","font-bold":""},V={"op-70":""},q=k({__name:"http-status-codes",setup($){const e=B(""),{searchResult:f}=b({search:e,data:i.flatMap(({codes:u,category:a})=>u.map(o=>({...o,category:a}))),options:{keys:[{name:"code",weight:3},{name:"name",weight:2},"description","category"]}}),v=F(()=>e.value?[{category:"Search results",codes:f.value}]:i);return(u,a)=>{const o=S,g=N;return t(),n("div",null,[w(o,{value:p(e),"onUpdate:value":a[0]||(a[0]=r=>C(e)?e.value=r:null),placeholder:"Search http status...",autofocus:"","raw-text":"","mb-10":""},null,8,["value"]),(t(!0),n(m,null,h(p(v),({codes:r,category:l})=>(t(),n("div",{key:l,"mb-8":""},[c("div",R,s(l),1),(t(!0),n(m,null,h(r,({code:_,description:y,name:x,type:d})=>(t(),z(g,{key:_,"mb-2":""},{default:E(()=>[c("div",T,s(_)+" "+s(x),1),c("div",V,s(y)+" "+s(d!=="HTTP"?`For ${d}.`:""),1)]),_:2},1024))),128))]))),128))])}}});export{q as default};
