import{d as q,z as u,E as B,w as t,y as U,o as $,k as o,n as l,B as p,q as F,m as L,C as N,bO as V,x as y,bL as z}from"./index-a828413e.js";import{_ as D}from"./c-select-7c7834cf.js";import{u as E}from"./useQRCode-ee05e817.js";import{u as R}from"./downloadBase64-52c89768.js";import{_ as T}from"./ColorPicker-b8641a76.js";import{_ as O}from"./FormItem-a888e5b4.js";import{_ as Q}from"./Form-0960ffb3.js";import{_ as Y}from"./Image-c8b18aed.js";import"./index-8ea3cd6a.js";import"./___vite-browser-external_commonjs-proxy-4d7ce06d.js";import"./Input-d144dbf3.js";import"./Button-0c407b94.js";import"./use-form-item-cff217fe.js";import"./use-locale-c394c8da.js";import"./InputGroup-55a483a4.js";const j={flex:"","flex-col":"","items-center":"","gap-3":""},_o=q({__name:"qr-code-generator",setup(A){const _=u("#000000ff"),a=u("#ffffffff"),r=u("medium"),f=["low","medium","quartile","high"],s=u("https://it-tools.tech"),{qrcode:i}=E({text:s,color:{background:a,foreground:_},errorCorrectionLevel:r,options:{width:1024}}),{download:g}=R({source:i,filename:"qr-code.png"});return(G,e)=>{const b=N,c=T,m=O,v=D,x=Q,d=V,w=Y,h=y,k=z,C=U;return $(),B(C,null,{default:t(()=>[o(k,{"x-gap":"12","y-gap":"12",cols:"1 600:3"},{default:t(()=>[o(d,{span:"2"},{default:t(()=>[o(b,{value:l(s),"onUpdate:value":e[0]||(e[0]=n=>p(s)?s.value=n:null),"label-position":"left","label-width":"130px","label-align":"right",label:"Text:",multiline:"",rows:"1",autosize:"",placeholder:"Your link or text...","mb-6":""},null,8,["value"]),o(x,{"label-width":"130","label-placement":"left"},{default:t(()=>[o(m,{label:"Foreground color:"},{default:t(()=>[o(c,{value:l(_),"onUpdate:value":e[1]||(e[1]=n=>p(_)?_.value=n:null),modes:["hex"]},null,8,["value"])]),_:1}),o(m,{label:"Background color:"},{default:t(()=>[o(c,{value:l(a),"onUpdate:value":e[2]||(e[2]=n=>p(a)?a.value=n:null),modes:["hex"]},null,8,["value"])]),_:1}),o(v,{value:l(r),"onUpdate:value":e[3]||(e[3]=n=>p(r)?r.value=n:null),label:"Error resistance:","label-position":"left","label-width":"130px","label-align":"right",options:f.map(n=>({label:n,value:n}))},null,8,["value","options"])]),_:1})]),_:1}),o(d,null,{default:t(()=>[F("div",j,[o(w,{src:l(i),width:"200"},null,8,["src"]),o(h,{onClick:l(g)},{default:t(()=>[L(" Download qr-code ")]),_:1},8,["onClick"])])]),_:1})]),_:1})]),_:1})}}});export{_o as default};