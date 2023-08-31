import{aQ as Be,aR as Me,b5 as Se,br as Te,a as De,c as W,d as Pe,u as Ce,b as re,z as R,a0 as Fe,a1 as Oe,S as g,ak as Ae,b6 as _e,e as Ue,h as s,ba as X,a_ as q,P as J,Q as Y,a6 as I,a5 as ke,bs as $e}from"./index-a828413e.js";import{u as Ee}from"./use-locale-c394c8da.js";import{u as Le}from"./use-form-item-cff217fe.js";import{N as He}from"./Input-d144dbf3.js";import{R as ze}from"./Remove-ddd6ebd8.js";import{X as Z}from"./Button-0c407b94.js";import{A as je}from"./Add-997e5c60.js";const Ge=n=>{const{textColorDisabled:i}=n;return{iconColorDisabled:i}},Ke=Be({name:"InputNumber",common:Me,peers:{Button:Se,Input:Te},self:Ge}),Qe=Ke;function We(n){return n==null||typeof n=="string"&&n.trim()===""?null:Number(n)}function Xe(n){return n.includes(".")&&(/^(-)?\d+.*(\.|0)$/.test(n)||/^\.\d+$/.test(n))}function L(n){return n==null?!0:!Number.isNaN(n)}function ee(n,i){return n==null?"":i===void 0?String(n):n.toFixed(i)}function H(n){if(n===null)return null;if(typeof n=="number")return n;{const i=Number(n);return Number.isNaN(i)?null:i}}const qe=De([W("input-number-suffix",`
 display: inline-block;
 margin-right: 10px;
 `),W("input-number-prefix",`
 display: inline-block;
 margin-left: 10px;
 `)]),ne=800,te=100,Je=Object.assign(Object.assign({},re.props),{autofocus:Boolean,loading:{type:Boolean,default:void 0},placeholder:String,defaultValue:{type:Number,default:null},value:Number,step:{type:[Number,String],default:1},min:[Number,String],max:[Number,String],size:String,disabled:{type:Boolean,default:void 0},validator:Function,bordered:{type:Boolean,default:void 0},showButton:{type:Boolean,default:!0},buttonPlacement:{type:String,default:"right"},readonly:Boolean,clearable:Boolean,keyboard:{type:Object,default:{}},updateValueOnInput:{type:Boolean,default:!0},parse:Function,format:Function,precision:Number,status:String,"onUpdate:value":[Function,Array],onUpdateValue:[Function,Array],onFocus:[Function,Array],onBlur:[Function,Array],onClear:[Function,Array],onChange:[Function,Array]}),un=Pe({name:"InputNumber",props:Je,setup(n){const{mergedBorderedRef:i,mergedClsPrefixRef:p,mergedRtlRef:D}=Ce(n),u=re("InputNumber","-input-number",qe,Qe,n,p),{localeRef:h}=Ee("InputNumber"),B=Le(n),{mergedSizeRef:le,mergedDisabledRef:ie,mergedStatusRef:ue}=B,c=R(null),z=R(null),j=R(null),P=R(n.defaultValue),ae=Fe(n,"value"),d=Oe(ae,P),v=R(""),C=e=>{const t=String(e).split(".")[1];return t?t.length:0},oe=e=>{const t=[n.min,n.max,n.step,e].map(r=>r===void 0?0:C(r));return Math.max(...t)},se=g(()=>{const{placeholder:e}=n;return e!==void 0?e:h.value.placeholder}),M=g(()=>{const e=H(n.step);return e!==null?e===0?1:Math.abs(e):1}),G=g(()=>{const e=H(n.min);return e!==null?e:null}),K=g(()=>{const e=H(n.max);return e!==null?e:null}),V=e=>{const{value:t}=d;if(e===t){b();return}const{"onUpdate:value":r,onUpdateValue:l,onChange:o}=n,{nTriggerFormInput:f,nTriggerFormChange:x}=B;o&&I(o,e),l&&I(l,e),r&&I(r,e),P.value=e,f(),x()},a=({offset:e,doUpdateIfValid:t,fixPrecision:r,isInputing:l})=>{const{value:o}=v;if(l&&Xe(o))return!1;const f=(n.parse||We)(o);if(f===null)return t&&V(null),null;if(L(f)){const x=C(f),{precision:N}=n;if(N!==void 0&&N<x&&!r)return!1;let m=parseFloat((f+e).toFixed(N??oe(f)));if(L(m)){const{value:$}=K,{value:E}=G;if($!==null&&m>$){if(!t||l)return!1;m=$}if(E!==null&&m<E){if(!t||l)return!1;m=E}return n.validator&&!n.validator(m)?!1:(t&&V(m),m)}}return!1},b=()=>{const{value:e}=d;if(L(e)){const{format:t,precision:r}=n;t?v.value=t(e):e===null||r===void 0||C(e)>r?v.value=ee(e,void 0):v.value=ee(e,r)}else v.value=String(e)};b();const de=g(()=>a({offset:0,doUpdateIfValid:!1,isInputing:!1,fixPrecision:!1})===!1),F=g(()=>{const{value:e}=d;if(n.validator&&e===null)return!1;const{value:t}=M;return a({offset:-t,doUpdateIfValid:!1,isInputing:!1,fixPrecision:!1})!==!1}),O=g(()=>{const{value:e}=d;if(n.validator&&e===null)return!1;const{value:t}=M;return a({offset:+t,doUpdateIfValid:!1,isInputing:!1,fixPrecision:!1})!==!1});function fe(e){const{onFocus:t}=n,{nTriggerFormFocus:r}=B;t&&I(t,e),r()}function ce(e){var t,r;if(e.target===((t=c.value)===null||t===void 0?void 0:t.wrapperElRef))return;const l=a({offset:0,doUpdateIfValid:!0,isInputing:!1,fixPrecision:!0});if(l!==!1){const x=(r=c.value)===null||r===void 0?void 0:r.inputElRef;x&&(x.value=String(l||"")),d.value===l&&b()}else b();const{onBlur:o}=n,{nTriggerFormBlur:f}=B;o&&I(o,e),f(),ke(()=>{b()})}function me(e){const{onClear:t}=n;t&&I(t,e)}function A(){const{value:e}=O;if(!e){k();return}const{value:t}=d;if(t===null)n.validator||V(Q());else{const{value:r}=M;a({offset:r,doUpdateIfValid:!0,isInputing:!1,fixPrecision:!0})}}function _(){const{value:e}=F;if(!e){U();return}const{value:t}=d;if(t===null)n.validator||V(Q());else{const{value:r}=M;a({offset:-r,doUpdateIfValid:!0,isInputing:!1,fixPrecision:!0})}}const he=fe,ve=ce;function Q(){if(n.validator)return null;const{value:e}=G,{value:t}=K;return e!==null?Math.max(0,e):t!==null?Math.min(0,t):0}function ge(e){me(e),V(null)}function pe(e){var t,r,l;!((t=j.value)===null||t===void 0)&&t.$el.contains(e.target)&&e.preventDefault(),!((r=z.value)===null||r===void 0)&&r.$el.contains(e.target)&&e.preventDefault(),(l=c.value)===null||l===void 0||l.activate()}let y=null,w=null,S=null;function U(){S&&(window.clearTimeout(S),S=null),y&&(window.clearInterval(y),y=null)}function k(){T&&(window.clearTimeout(T),T=null),w&&(window.clearInterval(w),w=null)}function be(){U(),S=window.setTimeout(()=>{y=window.setInterval(()=>{_()},te)},ne),q("mouseup",document,U,{once:!0})}let T=null;function xe(){k(),T=window.setTimeout(()=>{w=window.setInterval(()=>{A()},te)},ne),q("mouseup",document,k,{once:!0})}const Ie=()=>{w||A()},Ve=()=>{y||_()};function ye(e){var t,r;if(e.key==="Enter"){if(e.target===((t=c.value)===null||t===void 0?void 0:t.wrapperElRef))return;a({offset:0,doUpdateIfValid:!0,isInputing:!1,fixPrecision:!0})!==!1&&((r=c.value)===null||r===void 0||r.deactivate())}else if(e.key==="ArrowUp"){if(!O.value||n.keyboard.ArrowUp===!1)return;e.preventDefault(),a({offset:0,doUpdateIfValid:!0,isInputing:!1,fixPrecision:!0})!==!1&&A()}else if(e.key==="ArrowDown"){if(!F.value||n.keyboard.ArrowDown===!1)return;e.preventDefault(),a({offset:0,doUpdateIfValid:!0,isInputing:!1,fixPrecision:!0})!==!1&&_()}}function we(e){v.value=e,n.updateValueOnInput&&!n.format&&!n.parse&&n.precision===void 0&&a({offset:0,doUpdateIfValid:!0,isInputing:!0,fixPrecision:!1})}Ae(d,()=>{b()});const Ne={focus:()=>{var e;return(e=c.value)===null||e===void 0?void 0:e.focus()},blur:()=>{var e;return(e=c.value)===null||e===void 0?void 0:e.blur()}},Re=_e("InputNumber",D,p);return Object.assign(Object.assign({},Ne),{rtlEnabled:Re,inputInstRef:c,minusButtonInstRef:z,addButtonInstRef:j,mergedClsPrefix:p,mergedBordered:i,uncontrolledValue:P,mergedValue:d,mergedPlaceholder:se,displayedValueInvalid:de,mergedSize:le,mergedDisabled:ie,displayedValue:v,addable:O,minusable:F,mergedStatus:ue,handleFocus:he,handleBlur:ve,handleClear:ge,handleMouseDown:pe,handleAddClick:Ie,handleMinusClick:Ve,handleAddMousedown:xe,handleMinusMousedown:be,handleKeyDown:ye,handleUpdateDisplayedValue:we,mergedTheme:u,inputThemeOverrides:{paddingSmall:"0 8px 0 10px",paddingMedium:"0 8px 0 12px",paddingLarge:"0 8px 0 14px"},buttonThemeOverrides:Ue(()=>{const{self:{iconColorDisabled:e}}=u.value,[t,r,l,o]=$e(e);return{textColorTextDisabled:`rgb(${t}, ${r}, ${l})`,opacityDisabled:`${o}`}})})},render(){const{mergedClsPrefix:n,$slots:i}=this,p=()=>s(Z,{text:!0,disabled:!this.minusable||this.mergedDisabled||this.readonly,focusable:!1,theme:this.mergedTheme.peers.Button,themeOverrides:this.mergedTheme.peerOverrides.Button,builtinThemeOverrides:this.buttonThemeOverrides,onClick:this.handleMinusClick,onMousedown:this.handleMinusMousedown,ref:"minusButtonInstRef"},{icon:()=>J(i["minus-icon"],()=>[s(Y,{clsPrefix:n},{default:()=>s(ze,null)})])}),D=()=>s(Z,{text:!0,disabled:!this.addable||this.mergedDisabled||this.readonly,focusable:!1,theme:this.mergedTheme.peers.Button,themeOverrides:this.mergedTheme.peerOverrides.Button,builtinThemeOverrides:this.buttonThemeOverrides,onClick:this.handleAddClick,onMousedown:this.handleAddMousedown,ref:"addButtonInstRef"},{icon:()=>J(i["add-icon"],()=>[s(Y,{clsPrefix:n},{default:()=>s(je,null)})])});return s("div",{class:[`${n}-input-number`,this.rtlEnabled&&`${n}-input-number--rtl`]},s(He,{ref:"inputInstRef",autofocus:this.autofocus,status:this.mergedStatus,bordered:this.mergedBordered,loading:this.loading,value:this.displayedValue,onUpdateValue:this.handleUpdateDisplayedValue,theme:this.mergedTheme.peers.Input,themeOverrides:this.mergedTheme.peerOverrides.Input,builtinThemeOverrides:this.inputThemeOverrides,size:this.mergedSize,placeholder:this.mergedPlaceholder,disabled:this.mergedDisabled,readonly:this.readonly,textDecoration:this.displayedValueInvalid?"line-through":void 0,onFocus:this.handleFocus,onBlur:this.handleBlur,onKeydown:this.handleKeyDown,onMousedown:this.handleMouseDown,onClear:this.handleClear,clearable:this.clearable,internalLoadingBeforeSuffix:!0},{prefix:()=>{var u;return this.showButton&&this.buttonPlacement==="both"?[p(),X(i.prefix,h=>h?s("span",{class:`${n}-input-number-prefix`},h):null)]:(u=i.prefix)===null||u===void 0?void 0:u.call(i)},suffix:()=>{var u;return this.showButton?[X(i.suffix,h=>h?s("span",{class:`${n}-input-number-suffix`},h):null),this.buttonPlacement==="right"?p():null,D()]:(u=i.suffix)===null||u===void 0?void 0:u.call(i)}}))}});export{un as _};
