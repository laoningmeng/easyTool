import{L as m,a2 as c,e as r,aN as g,K as F}from"./index-a828413e.js";const a=F("n-form-item");function I(t,{defaultSize:s="medium",mergedSize:o,mergedDisabled:i}={}){const e=m(a,null);c(a,null);const f=r(o?()=>o(e):()=>{const{size:n}=t;if(n)return n;if(e){const{mergedSize:u}=e;if(u.value!==void 0)return u.value}return s}),d=r(i?()=>i(e):()=>{const{disabled:n}=t;return n!==void 0?n:e?e.disabled.value:!1}),l=r(()=>{const{status:n}=t;return n||(e==null?void 0:e.mergedValidationStatus.value)});return g(()=>{e&&e.restoreValidation()}),{mergedSizeRef:f,mergedDisabledRef:d,mergedStatusRef:l,nTriggerFormBlur(){e&&e.handleContentBlur()},nTriggerFormChange(){e&&e.handleContentChange()},nTriggerFormFocus(){e&&e.handleContentFocus()},nTriggerFormInput(){e&&e.handleContentInput()}}}export{a as f,I as u};
