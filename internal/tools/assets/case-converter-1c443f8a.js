import{d as P,z as R,e as E,E as p,w as T,y as U,o as C,k,bD as b,n as g,B as y,j as B,r as F,F as Z,q as j,C as z}from"./index-a828413e.js";import{_ as D}from"./InputCopyable.vue_vue_type_script_setup_true_lang-c867c80c.js";import"./content-copy-0fe7477a.js";import"./copy-1dd92e81.js";var l=function(){return l=Object.assign||function(e){for(var r,t=1,n=arguments.length;t<n;t++){r=arguments[t];for(var i in r)Object.prototype.hasOwnProperty.call(r,i)&&(e[i]=r[i])}return e},l.apply(this,arguments)};function S(a){return a.toLowerCase()}var $=[/([a-z0-9])([A-Z])/g,/([A-Z])([A-Z][a-z])/g],N=/[^A-Z0-9]+/gi;function u(a,e){e===void 0&&(e={});for(var r=e.splitRegexp,t=r===void 0?$:r,n=e.stripRegexp,i=n===void 0?N:n,c=e.transform,m=c===void 0?S:c,o=e.delimiter,s=o===void 0?" ":o,f=h(h(a,t,"$1\0$2"),i,"\0"),d=0,_=f.length;f.charAt(d)==="\0";)d++;for(;f.charAt(_-1)==="\0";)_--;return f.slice(d,_).split("\0").map(m).join(s)}function h(a,e,r){return e instanceof RegExp?a.replace(e,r):e.reduce(function(t,n){return t.replace(n,r)},a)}function w(a,e){var r=a.charAt(0),t=a.substr(1).toLowerCase();return e>0&&r>="0"&&r<="9"?"_"+r+t:""+r.toUpperCase()+t}function L(a,e){return e===void 0&&(e={}),u(a,l({delimiter:"",transform:w},e))}function O(a,e){return e===0?a.toLowerCase():w(a,e)}function G(a,e){return e===void 0&&(e={}),L(a,l({transform:O},e))}function A(a){return a.charAt(0).toUpperCase()+a.substr(1)}function I(a){return A(a.toLowerCase())}function x(a,e){return e===void 0&&(e={}),u(a,l({delimiter:" ",transform:I},e))}function V(a){return a.toUpperCase()}function X(a,e){return e===void 0&&(e={}),u(a,l({delimiter:"_",transform:V},e))}function v(a,e){return e===void 0&&(e={}),u(a,l({delimiter:"."},e))}function Y(a,e){return e===void 0&&(e={}),x(a,l({delimiter:"-"},e))}function q(a,e){return e===void 0&&(e={}),v(a,l({delimiter:"-"},e))}function H(a,e){return e===void 0&&(e={}),v(a,l({delimiter:"/"},e))}function W(a,e){var r=a.toLowerCase();return e===0?A(r):r}function J(a,e){return e===void 0&&(e={}),u(a,l({delimiter:" ",transform:W},e))}function K(a,e){return e===void 0&&(e={}),v(a,l({delimiter:"_"},e))}const M=j("div",{"my-16px":"",divider:""},null,-1),te=P({__name:"case-converter",setup(a){const e={stripRegexp:/[^A-Za-zÀ-ÖØ-öø-ÿ]+/gi},r=R("lorem ipsum dolor sit amet"),t=E(()=>[{label:"Lowercase:",value:u(r.value,e).toLocaleLowerCase()},{label:"Uppercase:",value:u(r.value,e).toLocaleUpperCase()},{label:"Camelcase:",value:G(r.value,e)},{label:"Capitalcase:",value:x(r.value,e)},{label:"Constantcase:",value:X(r.value,e)},{label:"Dotcase:",value:v(r.value,e)},{label:"Headercase:",value:Y(r.value,e)},{label:"Nocase:",value:u(r.value,e)},{label:"Paramcase:",value:q(r.value,e)},{label:"Pascalcase:",value:L(r.value,e)},{label:"Pathcase:",value:H(r.value,e)},{label:"Sentencecase:",value:J(r.value,e)},{label:"Snakecase:",value:K(r.value,e)}]),n={labelPosition:"left",labelWidth:"120px",labelAlign:"right"};return(i,c)=>{const m=z,o=U;return C(),p(o,null,{default:T(()=>[k(m,b({value:g(r),"onUpdate:value":c[0]||(c[0]=s=>y(r)?r.value=s:null),label:"Your string:",placeholder:"Your string...","raw-text":""},n),null,16,["value"]),M,(C(!0),B(Z,null,F(g(t),s=>(C(),p(D,b({key:s.label,value:s.value,label:s.label},n,{"mb-1":""}),null,16,["value","label"]))),128))]),_:1})}}});export{te as default};