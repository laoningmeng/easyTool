import{d as V,z as m,e as _,A as v,j as Y,k as o,w as i,F as k,y as z,o as E,n as t,B as f,q as g,m as b,C as R,x as B}from"./index-a828413e.js";import{u as y}from"./copy-1dd92e81.js";import{i as C}from"./boolean-c7e7c785.js";import{w}from"./defaults-4d6daddf.js";const D={flex:"","justify-center":""},N={flex:"","justify-center":""},F=V({__name:"url-encoder",setup(T){const n=m("Hello world :)"),r=_(()=>w(()=>encodeURIComponent(n.value),"")),x=v({source:n,rules:[{validator:a=>C(()=>encodeURIComponent(a)),message:"Impossible to parse this string"}]}),{copy:h}=y({source:r,text:"Encoded string copied to the clipboard"}),s=m("Hello%20world%20%3A)"),u=_(()=>w(()=>decodeURIComponent(s.value),"")),I=v({source:n,rules:[{validator:a=>C(()=>decodeURIComponent(a)),message:"Impossible to parse this string"}]}),{copy:U}=y({source:u,text:"Decoded string copied to the clipboard"});return(a,e)=>{const d=R,c=B,p=z;return E(),Y(k,null,[o(p,{title:"Encode"},{default:i(()=>[o(d,{value:t(n),"onUpdate:value":e[0]||(e[0]=l=>f(n)?n.value=l:null),label:"Your string :",validation:t(x),multiline:"",autosize:"",placeholder:"The string to encode",rows:"2","mb-3":""},null,8,["value","validation"]),o(d,{label:"Your string encoded :",value:t(r),multiline:"",autosize:"",readonly:"",placeholder:"Your string encoded",rows:"2","mb-3":""},null,8,["value"]),g("div",D,[o(c,{onClick:e[1]||(e[1]=l=>t(h)())},{default:i(()=>[b(" Copy ")]),_:1})])]),_:1}),o(p,{title:"Decode"},{default:i(()=>[o(d,{value:t(s),"onUpdate:value":e[2]||(e[2]=l=>f(s)?s.value=l:null),label:"Your encoded string :",validation:t(I),multiline:"",autosize:"",placeholder:"The string to decode",rows:"2","mb-3":""},null,8,["value","validation"]),o(d,{label:"Your string decoded :",value:t(u),multiline:"",autosize:"",readonly:"",placeholder:"Your string decoded",rows:"2","mb-3":""},null,8,["value"]),g("div",N,[o(c,{onClick:e[3]||(e[3]=l=>t(U)())},{default:i(()=>[b(" Copy ")]),_:1})])]),_:1})],64)}}});export{F as default};
