import{_ as i}from"./FormatTransformer.vue_vue_type_script_setup_true_lang-f773df0a.js";import{s as r}from"./public-api-aa5b1ac9.js";import{J as t}from"./index-d4e515d0.js";import{i as m}from"./boolean-c7e7c785.js";import{w as n}from"./defaults-4d6daddf.js";import{d as p,E as l,o as u}from"./index-a828413e.js";import"./TextareaCopyable-76ae05c4.js";import"./copy-1dd92e81.js";import"./Copy-671934bd.js";import"./Scrollbar-ae9c5e6e.js";const w=p({__name:"json-to-yaml",setup(f){const a=o=>n(()=>r(t.parse(o)),""),e=[{validator:o=>o===""||m(()=>r(t.parse(o))),message:"Provided JSON is not valid."}];return(o,c)=>{const s=i;return u(),l(s,{"input-label":"Your JSON","input-placeholder":"Paste your JSON here...","output-label":"YAML from your JSON","output-language":"yaml","input-validation-rules":e,transformer:a})}}});export{w as default};
