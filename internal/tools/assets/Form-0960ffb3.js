import{c as b,W as I,a as w,d as v,u as _,b as x,z as j,a2 as g,h as M,bo as S,b$ as y}from"./index-a828413e.js";import{f as O,a as P}from"./FormItem-a888e5b4.js";const B=b("form",[I("inline",`
 width: 100%;
 display: inline-flex;
 align-items: flex-start;
 align-content: space-around;
 `,[b("form-item",{width:"auto",marginRight:"18px"},[w("&:last-child",{marginRight:0})])])]);var C=globalThis&&globalThis.__awaiter||function(n,c,o,s){function u(r){return r instanceof o?r:new o(function(a){a(r)})}return new(o||(o=Promise))(function(r,a){function h(t){try{i(s.next(t))}catch(m){a(m)}}function e(t){try{i(s.throw(t))}catch(m){a(m)}}function i(t){t.done?r(t.value):u(t.value).then(h,e)}i((s=s.apply(n,c||[])).next())})};const L=Object.assign(Object.assign({},x.props),{inline:Boolean,labelWidth:[Number,String],labelAlign:String,labelPlacement:{type:String,default:"top"},model:{type:Object,default:()=>{}},rules:Object,disabled:Boolean,size:String,showRequireMark:{type:Boolean,default:void 0},requireMarkPlacement:String,showFeedback:{type:Boolean,default:!0},onSubmit:{type:Function,default:n=>n.preventDefault()},showLabel:{type:Boolean,default:void 0},validateMessages:Object}),V=v({name:"Form",props:L,setup(n){const{mergedClsPrefixRef:c}=_(n);x("Form","-form",B,S,n,c);const o={},s=j(void 0),u=e=>{const i=s.value;(i===void 0||e>=i)&&(s.value=e)};function r(e,i=()=>!0){return C(this,void 0,void 0,function*(){return yield new Promise((t,m)=>{const p=[];for(const d of y(o)){const f=o[d];for(const l of f)l.path&&p.push(l.internalValidate(null,i))}Promise.all(p).then(d=>{if(d.some(f=>!f.valid)){const f=d.filter(l=>l.errors).map(l=>l.errors);e&&e(f),m(f)}else e&&e(),t()})})})}function a(){for(const e of y(o)){const i=o[e];for(const t of i)t.restoreValidation()}}return g(O,{props:n,maxChildLabelWidthRef:s,deriveMaxChildLabelWidth:u}),g(P,{formItems:o}),Object.assign({validate:r,restoreValidation:a},{mergedClsPrefix:c})},render(){const{mergedClsPrefix:n}=this;return M("form",{class:[`${n}-form`,this.inline&&`${n}-form--inline`],onSubmit:this.onSubmit},this.$slots)}});export{V as _};
