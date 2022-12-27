/*! For license information please see cab19fec.chunk.js.LICENSE.txt */
(self.webpackChunkw3social_interface=self.webpackChunkw3social_interface||[]).push([[844],{7410:(e,o,t)=>{"use strict";var r=t(489);o.Z=void 0;var a=r(t(5872)),n=t(1527),i=(0,a.default)((0,n.jsx)("path",{d:"M12 22c1.1 0 2-.9 2-2h-4c0 1.1.9 2 2 2zm6-6v-5c0-3.07-1.63-5.64-4.5-6.32V4c0-.83-.67-1.5-1.5-1.5s-1.5.67-1.5 1.5v.68C7.64 5.36 6 7.92 6 11v5l-2 2v1h16v-1l-2-2zm-2 1H8v-6c0-2.48 1.51-4.5 4-4.5s4 2.02 4 4.5v6z"}),"NotificationsOutlined");o.Z=i},1474:(e,o,t)=>{"use strict";t.d(o,{Z:()=>Z});var r=t(4280),a=t(7161),n=t(959),i=t(5924),l=t(945),s=t(8414),c=t(1471),d=t(7683),p=t(1138),u=t(1336),v=t(4379),g=t(83);function f(e){return(0,g.Z)("MuiIconButton",e)}const m=(0,v.Z)("MuiIconButton",["root","disabled","colorInherit","colorPrimary","colorSecondary","colorError","colorInfo","colorSuccess","colorWarning","edgeStart","edgeEnd","sizeSmall","sizeMedium","sizeLarge"]);var b=t(1527);const h=["edge","children","className","color","disabled","disableFocusRipple","size"],x=(0,c.ZP)(p.Z,{name:"MuiIconButton",slot:"Root",overridesResolver:(e,o)=>{const{ownerState:t}=e;return[o.root,"default"!==t.color&&o[`color${(0,u.Z)(t.color)}`],t.edge&&o[`edge${(0,u.Z)(t.edge)}`],o[`size${(0,u.Z)(t.size)}`]]}})((({theme:e,ownerState:o})=>(0,a.Z)({textAlign:"center",flex:"0 0 auto",fontSize:e.typography.pxToRem(24),padding:8,borderRadius:"50%",overflow:"visible",color:(e.vars||e).palette.action.active,transition:e.transitions.create("background-color",{duration:e.transitions.duration.shortest})},!o.disableRipple&&{"&:hover":{backgroundColor:e.vars?`rgba(${e.vars.palette.action.activeChannel} / ${e.vars.palette.action.hoverOpacity})`:(0,s.Fq)(e.palette.action.active,e.palette.action.hoverOpacity),"@media (hover: none)":{backgroundColor:"transparent"}}},"start"===o.edge&&{marginLeft:"small"===o.size?-3:-12},"end"===o.edge&&{marginRight:"small"===o.size?-3:-12})),(({theme:e,ownerState:o})=>{var t;const r=null==(t=(e.vars||e).palette)?void 0:t[o.color];return(0,a.Z)({},"inherit"===o.color&&{color:"inherit"},"inherit"!==o.color&&"default"!==o.color&&(0,a.Z)({color:null==r?void 0:r.main},!o.disableRipple&&{"&:hover":(0,a.Z)({},r&&{backgroundColor:e.vars?`rgba(${r.mainChannel} / ${e.vars.palette.action.hoverOpacity})`:(0,s.Fq)(r.main,e.palette.action.hoverOpacity)},{"@media (hover: none)":{backgroundColor:"transparent"}})}),"small"===o.size&&{padding:5,fontSize:e.typography.pxToRem(18)},"large"===o.size&&{padding:12,fontSize:e.typography.pxToRem(28)},{[`&.${m.disabled}`]:{backgroundColor:"transparent",color:(e.vars||e).palette.action.disabled}})})),Z=n.forwardRef((function(e,o){const t=(0,d.Z)({props:e,name:"MuiIconButton"}),{edge:n=!1,children:s,className:c,color:p="default",disabled:v=!1,disableFocusRipple:g=!1,size:m="medium"}=t,Z=(0,r.Z)(t,h),k=(0,a.Z)({},t,{edge:n,color:p,disabled:v,disableFocusRipple:g,size:m}),w=(e=>{const{classes:o,disabled:t,color:r,edge:a,size:n}=e,i={root:["root",t&&"disabled","default"!==r&&`color${(0,u.Z)(r)}`,a&&`edge${(0,u.Z)(a)}`,`size${(0,u.Z)(n)}`]};return(0,l.Z)(i,f,o)})(k);return(0,b.jsx)(x,(0,a.Z)({className:(0,i.Z)(w.root,c),centerRipple:!0,focusRipple:!g,disabled:v,ref:o,ownerState:k},Z,{children:s}))}))},2868:(e,o,t)=>{"use strict";t.d(o,{Z:()=>h});var r=t(4280),a=t(7161),n=t(959),i=t(5924),l=t(945),s=t(8414),c=t(1471);const d=e=>{let o;return o=e<1?5.11916*e**2:4.5*Math.log(e+1)+2,(o/100).toFixed(2)};var p=t(7683),u=t(4379),v=t(83);function g(e){return(0,v.Z)("MuiPaper",e)}(0,u.Z)("MuiPaper",["root","rounded","outlined","elevation","elevation0","elevation1","elevation2","elevation3","elevation4","elevation5","elevation6","elevation7","elevation8","elevation9","elevation10","elevation11","elevation12","elevation13","elevation14","elevation15","elevation16","elevation17","elevation18","elevation19","elevation20","elevation21","elevation22","elevation23","elevation24"]);var f=t(1527);const m=["className","component","elevation","square","variant"],b=(0,c.ZP)("div",{name:"MuiPaper",slot:"Root",overridesResolver:(e,o)=>{const{ownerState:t}=e;return[o.root,o[t.variant],!t.square&&o.rounded,"elevation"===t.variant&&o[`elevation${t.elevation}`]]}})((({theme:e,ownerState:o})=>{var t;return(0,a.Z)({backgroundColor:(e.vars||e).palette.background.paper,color:(e.vars||e).palette.text.primary,transition:e.transitions.create("box-shadow")},!o.square&&{borderRadius:e.shape.borderRadius},"outlined"===o.variant&&{border:`1px solid ${(e.vars||e).palette.divider}`},"elevation"===o.variant&&(0,a.Z)({boxShadow:(e.vars||e).shadows[o.elevation]},!e.vars&&"dark"===e.palette.mode&&{backgroundImage:`linear-gradient(${(0,s.Fq)("#fff",d(o.elevation))}, ${(0,s.Fq)("#fff",d(o.elevation))})`},e.vars&&{backgroundImage:null==(t=e.vars.overlays)?void 0:t[o.elevation]}))})),h=n.forwardRef((function(e,o){const t=(0,p.Z)({props:e,name:"MuiPaper"}),{className:n,component:s="div",elevation:c=1,square:d=!1,variant:u="elevation"}=t,v=(0,r.Z)(t,m),h=(0,a.Z)({},t,{component:s,elevation:c,square:d,variant:u}),x=(e=>{const{square:o,elevation:t,variant:r,classes:a}=e,n={root:["root",r,!o&&"rounded","elevation"===r&&`elevation${t}`]};return(0,l.Z)(n,g,a)})(h);return(0,f.jsx)(b,(0,a.Z)({as:s,ownerState:h,className:(0,i.Z)(x.root,n),ref:o},v))}))},4875:(e,o)=>{var t;!function(){"use strict";var r={}.hasOwnProperty;function a(){for(var e=[],o=0;o<arguments.length;o++){var t=arguments[o];if(t){var n=typeof t;if("string"===n||"number"===n)e.push(t);else if(Array.isArray(t)){if(t.length){var i=a.apply(null,t);i&&e.push(i)}}else if("object"===n){if(t.toString!==Object.prototype.toString&&!t.toString.toString().includes("[native code]")){e.push(t.toString());continue}for(var l in t)r.call(t,l)&&t[l]&&e.push(l)}}}return e.join(" ")}e.exports?(a.default=a,e.exports=a):void 0===(t=function(){return a}.apply(o,[]))||(e.exports=t)}()},259:(e,o,t)=>{"use strict";t.r(o),t.d(o,{default:()=>O});var r=t(1527),a=t(7626),n=t(959),i=t(4280),l=t(7161),s=t(5924),c=t(945),d=t(1471),p=t(7683),u=t(1336),v=t(2868),g=t(4379),f=t(83);function m(e){return(0,f.Z)("MuiAppBar",e)}(0,g.Z)("MuiAppBar",["root","positionFixed","positionAbsolute","positionSticky","positionStatic","positionRelative","colorDefault","colorPrimary","colorSecondary","colorInherit","colorTransparent"]);const b=["className","color","enableColorOnDark","position"],h=(e,o)=>`${null==e?void 0:e.replace(")","")}, ${o})`,x=(0,d.ZP)(v.Z,{name:"MuiAppBar",slot:"Root",overridesResolver:(e,o)=>{const{ownerState:t}=e;return[o.root,o[`position${(0,u.Z)(t.position)}`],o[`color${(0,u.Z)(t.color)}`]]}})((({theme:e,ownerState:o})=>{const t="light"===e.palette.mode?e.palette.grey[100]:e.palette.grey[900];return(0,l.Z)({display:"flex",flexDirection:"column",width:"100%",boxSizing:"border-box",flexShrink:0},"fixed"===o.position&&{position:"fixed",zIndex:(e.vars||e).zIndex.appBar,top:0,left:"auto",right:0,"@media print":{position:"absolute"}},"absolute"===o.position&&{position:"absolute",zIndex:(e.vars||e).zIndex.appBar,top:0,left:"auto",right:0},"sticky"===o.position&&{position:"sticky",zIndex:(e.vars||e).zIndex.appBar,top:0,left:"auto",right:0},"static"===o.position&&{position:"static"},"relative"===o.position&&{position:"relative"},!e.vars&&(0,l.Z)({},"default"===o.color&&{backgroundColor:t,color:e.palette.getContrastText(t)},o.color&&"default"!==o.color&&"inherit"!==o.color&&"transparent"!==o.color&&{backgroundColor:e.palette[o.color].main,color:e.palette[o.color].contrastText},"inherit"===o.color&&{color:"inherit"},"dark"===e.palette.mode&&!o.enableColorOnDark&&{backgroundColor:null,color:null},"transparent"===o.color&&(0,l.Z)({backgroundColor:"transparent",color:"inherit"},"dark"===e.palette.mode&&{backgroundImage:"none"})),e.vars&&(0,l.Z)({},"default"===o.color&&{"--AppBar-background":o.enableColorOnDark?e.vars.palette.AppBar.defaultBg:h(e.vars.palette.AppBar.darkBg,e.vars.palette.AppBar.defaultBg),"--AppBar-color":o.enableColorOnDark?e.vars.palette.text.primary:h(e.vars.palette.AppBar.darkColor,e.vars.palette.text.primary)},o.color&&!o.color.match(/^(default|inherit|transparent)$/)&&{"--AppBar-background":o.enableColorOnDark?e.vars.palette[o.color].main:h(e.vars.palette.AppBar.darkBg,e.vars.palette[o.color].main),"--AppBar-color":o.enableColorOnDark?e.vars.palette[o.color].contrastText:h(e.vars.palette.AppBar.darkColor,e.vars.palette[o.color].contrastText)},{backgroundColor:"var(--AppBar-background)",color:"inherit"===o.color?"inherit":"var(--AppBar-color)"},"transparent"===o.color&&{backgroundImage:"none",backgroundColor:"transparent",color:"inherit"}))})),Z=n.forwardRef((function(e,o){const t=(0,p.Z)({props:e,name:"MuiAppBar"}),{className:a,color:n="primary",enableColorOnDark:d=!1,position:v="fixed"}=t,g=(0,i.Z)(t,b),f=(0,l.Z)({},t,{color:n,position:v,enableColorOnDark:d}),h=(e=>{const{color:o,position:t,classes:r}=e,a={root:["root",`color${(0,u.Z)(o)}`,`position${(0,u.Z)(t)}`]};return(0,c.Z)(a,m,r)})(f);return(0,r.jsx)(x,(0,l.Z)({square:!0,component:"header",ownerState:f,elevation:4,className:(0,s.Z)(h.root,a,"fixed"===v&&"mui-fixed"),ref:o},g))}));var k=t(2513);function w(e){return(0,f.Z)("MuiToolbar",e)}(0,g.Z)("MuiToolbar",["root","gutters","regular","dense"]);const y=["className","component","disableGutters","variant"],S=(0,d.ZP)("div",{name:"MuiToolbar",slot:"Root",overridesResolver:(e,o)=>{const{ownerState:t}=e;return[o.root,!t.disableGutters&&o.gutters,o[t.variant]]}})((({theme:e,ownerState:o})=>(0,l.Z)({position:"relative",display:"flex",alignItems:"center"},!o.disableGutters&&{paddingLeft:e.spacing(2),paddingRight:e.spacing(2),[e.breakpoints.up("sm")]:{paddingLeft:e.spacing(3),paddingRight:e.spacing(3)}},"dense"===o.variant&&{minHeight:48})),(({theme:e,ownerState:o})=>"regular"===o.variant&&e.mixins.toolbar)),C=n.forwardRef((function(e,o){const t=(0,p.Z)({props:e,name:"MuiToolbar"}),{className:a,component:n="div",disableGutters:d=!1,variant:u="regular"}=t,v=(0,i.Z)(t,y),g=(0,l.Z)({},t,{component:n,disableGutters:d,variant:u}),f=(e=>{const{classes:o,disableGutters:t,variant:r}=e,a={root:["root",!t&&"gutters",r]};return(0,c.Z)(a,w,o)})(g);return(0,r.jsx)(S,(0,l.Z)({as:n,className:(0,s.Z)(f.root,a),ref:o,ownerState:g},v))}));var B=t(1474),R=t(7410),z=t(8642),j=t(4875),N=t.n(j);const $=function(){var e,o=(0,a.s0)(),t=(0,n.useState)((null==(e=window.location.pathname)?void 0:e.includes("square"))?1:0),r=t[0],i=t[1],l=(0,n.useState)(!0),s=l[0],c=(l[1],["","square"]);return{activeBtnIndex:r,needNotify:s,buttonSwitch:function(e){i(e),o(c[e],{replace:!0})},getButtonClass:function(e){var o="flex-1 rounded-full m-[5px] h-[25px] shadow-none";return r===e?N()(o,"bg-black text-white transition-colors"):N()(o,"bg-transparent text-black")}}},A=t.p+"assets/94f0e603.png";function I(){var e=$(),o=e.needNotify,t=e.getButtonClass,n=e.buttonSwitch,i=(0,a.s0)();return(0,r.jsx)(Z,{className:"bg-transparent shadow-none",position:"static",children:(0,r.jsxs)(C,{className:"flex px-0 box-border h-[35px]",children:[(0,r.jsx)(B.Z,{onClick:function(){return i("/aUser")},edge:"start",className:"text-black flex-1 mx-0 text-[30px]","aria-label":"menu",children:(0,r.jsx)("img",{src:A,alt:"myicon",className:"w-[30px] h-[30px]"})}),(0,r.jsxs)("div",{className:"flex-6 flex box-border rounded-full bg-white",children:[(0,r.jsx)(k.Z,{className:t(0),onClick:function(){return n(0)},variant:"contained",children:"DID 推送"}),(0,r.jsx)(k.Z,{className:t(1),onClick:function(){return n(1)},variant:"contained",children:"需求广场"})]}),(0,r.jsx)(B.Z,{edge:"start",className:"text-black flex-1 mx-0 text-[30px] pr-0","aria-label":"menu",onClick:function(){return i("/message")},children:o?(0,r.jsx)(z.Z,{color:"error",variant:"dot",children:(0,r.jsx)(R.Z,{})}):(0,r.jsx)(R.Z,{})})]})})}const M=function(e){var o=e.children;return(0,r.jsx)("div",{className:"p-[16px] relative flex-1",children:o})};const O=function(){return(0,r.jsxs)("div",{className:"bg-gray-100 box-border h-full flex flex-col",children:[(0,r.jsx)(I,{}),(0,r.jsx)(M,{children:(0,r.jsx)(a.j3,{})})]})}}}]);