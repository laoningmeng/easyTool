import{ah as m,d as J,z as k,ab as f,e as O,A as S,j as N,q as w,k as n,w as p,n as s,F as j,o as V,B as c,C as z,p as C}from"./index-a828413e.js";import{J as y}from"./index-d4e515d0.js";import{w as B}from"./defaults-4d6daddf.js";import{T as E}from"./TextareaCopyable-76ae05c4.js";import{_ as A}from"./FormItem-a888e5b4.js";import{_ as P}from"./Switch-72472d29.js";import{_ as U}from"./InputNumber-a6e19667.js";import"./copy-1dd92e81.js";import"./Copy-671934bd.js";import"./Scrollbar-ae9c5e6e.js";import"./use-form-item-cff217fe.js";import"./use-locale-c394c8da.js";import"./Input-d144dbf3.js";import"./Button-0c407b94.js";import"./Remove-ddd6ebd8.js";import"./Add-997e5c60.js";function d(t){return typeof t!="object"||t===null?t:Array.isArray(t)?t.map(d):Object.keys(t).sort((o,e)=>o.localeCompare(e)).reduce((o,e)=>(o[e]=d(t[e]),o),{})}function F({rawJson:t,sortKeys:o=!0,indentSize:e=3}){const a=y.parse(m(t));return JSON.stringify(m(o)?d(a):a,null,m(e))}const I={style:{flex:"0 0 100%"}},K={style:{margin:"0 auto","max-width":"600px"},flex:"","justify-center":"","gap-3":""},T=J({__name:"json-viewer",setup(t){const o=k(),e=f("json-prettify:raw-json",'{"hello": "world", "foo": "bar"}'),a=f("json-prettify:indent-size",3),i=f("json-prettify:sort-keys",!0),g=O(()=>B(()=>F({rawJson:e,indentSize:a,sortKeys:i}),"")),v=S({source:e,rules:[{validator:_=>_===""||y.parse(_),message:"Provided JSON is not valid."}]});return(_,l)=>{const b=P,u=A,h=U,x=z;return V(),N(j,null,[w("div",I,[w("div",K,[n(u,{label:"Sort keys :","label-placement":"left","label-width":"100"},{default:p(()=>[n(b,{value:s(i),"onUpdate:value":l[0]||(l[0]=r=>c(i)?i.value=r:null)},null,8,["value"])]),_:1}),n(u,{label:"Indent size :","label-placement":"left","label-width":"100","show-feedback":!1},{default:p(()=>[n(h,{value:s(a),"onUpdate:value":l[1]||(l[1]=r=>c(a)?a.value=r:null),min:"0",max:"10",style:{width:"100px"}},null,8,["value"])]),_:1})])]),n(u,{label:"Your raw JSON",feedback:s(v).message,"validation-status":s(v).status},{default:p(()=>[n(x,{ref_key:"inputElement",ref:o,value:s(e),"onUpdate:value":l[2]||(l[2]=r=>c(e)?e.value=r:null),placeholder:"Paste your raw JSON here...",rows:"20",multiline:"",autocomplete:"off",autocorrect:"off",autocapitalize:"off",spellcheck:"false",monospace:""},null,8,["value"])]),_:1},8,["feedback","validation-status"]),n(u,{label:"Prettified version of your JSON"},{default:p(()=>[n(E,{value:s(g),language:"json","follow-height-of":s(o)},null,8,["value","follow-height-of"])]),_:1})],64)}}});const ne=C(T,[["__scopeId","data-v-2c415be6"]]);export{ne as default};
