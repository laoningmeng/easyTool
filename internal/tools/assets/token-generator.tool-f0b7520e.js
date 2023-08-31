import{d as U,cL as V,j as N,k as e,w as n,y as T,o as $,q as v,n as t,B as a,m as k,C as B,x as L,p as j}from"./index-a828413e.js";import{c as R}from"./token-generator.service-0ba2a8dc.js";import{u as G}from"./copy-1dd92e81.js";import{u as r}from"./queryParams-9fa1d322.js";import{a as I}from"./computedRefreshable-05f72a99.js";import{_ as S}from"./Switch-72472d29.js";import{_ as q}from"./FormItem-a888e5b4.js";import{_ as z}from"./Form-0960ffb3.js";import{_ as E}from"./Slider-c3aac5cb.js";import"./use-form-item-cff217fe.js";const P={flex:"","justify-center":""},Q={"mt-5":"",flex:"","justify-center":"","gap-3":""},A=U({__name:"token-generator.tool",setup(D){const u=r({name:"length",defaultValue:64}),_=r({name:"uppercase",defaultValue:!0}),p=r({name:"lowercase",defaultValue:!0}),m=r({name:"numbers",defaultValue:!0}),i=r({name:"symbols",defaultValue:!1}),{t:c}=V(),[d,g]=I(()=>R({length:u.value,withUppercase:_.value,withLowercase:p.value,withNumbers:m.value,withSymbols:i.value})),{copy:w}=G({source:d,text:"Token copied to the clipboard"});return(F,l)=>{const f=S,s=q,y=z,x=E,h=B,b=L,C=T;return $(),N("div",null,[e(C,null,{default:n(()=>[e(y,{"label-placement":"left","label-width":"140"},{default:n(()=>[v("div",P,[v("div",null,[e(s,{label:t(c)("tools.token-generator.uppercase")},{default:n(()=>[e(f,{value:t(_),"onUpdate:value":l[0]||(l[0]=o=>a(_)?_.value=o:null)},null,8,["value"])]),_:1},8,["label"]),e(s,{label:t(c)("tools.token-generator.lowercase")},{default:n(()=>[e(f,{value:t(p),"onUpdate:value":l[1]||(l[1]=o=>a(p)?p.value=o:null)},null,8,["value"])]),_:1},8,["label"])]),v("div",null,[e(s,{label:t(c)("tools.token-generator.numbers")},{default:n(()=>[e(f,{value:t(m),"onUpdate:value":l[2]||(l[2]=o=>a(m)?m.value=o:null)},null,8,["value"])]),_:1},8,["label"]),e(s,{label:t(c)("tools.token-generator.symbols")},{default:n(()=>[e(f,{value:t(i),"onUpdate:value":l[3]||(l[3]=o=>a(i)?i.value=o:null)},null,8,["value"])]),_:1},8,["label"])])])]),_:1}),e(s,{label:`Length (${t(u)})`,"label-placement":"left"},{default:n(()=>[e(x,{value:t(u),"onUpdate:value":l[4]||(l[4]=o=>a(u)?u.value=o:null),step:1,min:1,max:512},null,8,["value"])]),_:1},8,["label"]),e(h,{value:t(d),"onUpdate:value":l[5]||(l[5]=o=>a(d)?d.value=o:null),multiline:"",placeholder:"The token...",readonly:"",rows:"3",autosize:"",class:"token-display"},null,8,["value"]),v("div",Q,[e(b,{onClick:l[6]||(l[6]=o=>t(w)())},{default:n(()=>[k(" Copy ")]),_:1}),e(b,{onClick:t(g)},{default:n(()=>[k(" Refresh ")]),_:1},8,["onClick"])])]),_:1})])}}});const le=j(A,[["__scopeId","data-v-8c832e3a"]]);export{le as default};