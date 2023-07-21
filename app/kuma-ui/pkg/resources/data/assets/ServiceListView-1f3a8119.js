import{d as $,u as q,q as s,o as T,a as w,w as u,h as c,b as h,g as z,P as A,f as B}from"./index-f4ec0be6.js";import{S as F}from"./ServiceSummary-0e1acf9d.js";import{g as M,i as O,A as Q,_ as R}from"./RouteView.vue_vue_type_script_setup_true_lang-09fd82a0.js";import{_ as U}from"./RouteTitle.vue_vue_type_script_setup_true_lang-d39e0ee5.js";import{C as G}from"./ContentWrapper-67b66133.js";import{D as K}from"./DataOverview-706f2010.js";import{Q as y}from"./QueryParameter-70743f73.js";import"./kongponents.es-fed304fd.js";import"./DefinitionListItem-0b3f80a7.js";import"./ResourceCodeBlock.vue_vue_type_script_setup_true_lang-7c8e38a7.js";import"./CodeBlock.vue_vue_type_style_index_0_lang-56242a2b.js";import"./TextWithCopyButton-ae3a8132.js";import"./StatusInfo.vue_vue_type_script_setup_true_lang-7a761e1d.js";import"./ErrorBlock-f3ed6714.js";import"./StatusBadge-0a8731e8.js";import"./TagList-0012d9cb.js";const ie=$({__name:"ServiceListView",props:{selectedServiceName:{type:String,required:!1,default:null},offset:{type:Number,required:!1,default:0}},setup(k){const m=k,p=M(),{t:D}=O(),N=[{label:"Name",key:"entity"},{label:"Type",key:"serviceType"},{label:"Address",key:"addressPort"},{label:"Status",key:"status"},{label:"DP proxies (online / total)",key:"dpProxiesStatus"}],P={title:"No Data",message:"There are no service insights present."},E=q(),v=s(!0),d=s(null),_=s(null),S=s(m.offset),t=s(null),o=s(null),n=s({headers:N,data:[]});b(m.offset);async function b(a){var i;S.value=a,y.set("offset",a>0?a:null),v.value=!0,d.value=null;const e=E.params.mesh,r=A;try{const{items:l,next:f}=await p.getAllServiceInsightsFromMesh({mesh:e},{size:r,offset:a});_.value=f,n.value.data=V(l??[]),await x({name:m.selectedServiceName??((i=n.value.data[0])==null?void 0:i.entity.name),mesh:e})}catch(l){n.value.data=[],t.value=null,o.value=null,l instanceof Error?d.value=l:console.error(l)}finally{v.value=!1}}function V(a){return a.map(e=>{const{serviceType:r="internal",addressPort:i="",status:l="not_available"}=e,f={name:"service-detail-view",params:{mesh:e.mesh,service:e.name}};let g="—";if(e.dataplanes){const{online:C=0,total:L=0}=e.dataplanes;g=`${C} / ${L}`}return{entity:e,detailViewRoute:f,status:l,serviceType:r,dpProxiesStatus:g,addressPort:i}})}async function x({name:a,mesh:e}){a!==void 0?(t.value=await p.getServiceInsight({mesh:e,name:a}),t.value.serviceType==="external"&&(o.value=await p.getExternalServiceByServiceInsightName(e,a)),y.set("service",a)):(t.value=null,o.value=null,y.set("service",null))}return(a,e)=>(T(),w(R,null,{default:u(()=>[c(U,{title:h(D)("services.routes.items.title")},null,8,["title"]),z(),c(Q,null,{default:u(()=>[c(G,null,{content:u(()=>{var r;return[c(K,{"selected-entity-name":(r=t.value)==null?void 0:r.name,"page-size":h(A),error:d.value,"is-loading":v.value,"empty-state":P,"table-data":n.value,"table-data-is-empty":n.value.data.length===0,next:_.value,"page-offset":S.value,onTableAction:x,onLoadData:b},null,8,["selected-entity-name","page-size","error","is-loading","table-data","table-data-is-empty","next","page-offset"])]}),sidebar:u(()=>[t.value!==null?(T(),w(F,{key:0,service:t.value,"external-service":o.value},null,8,["service","external-service"])):B("",!0)]),_:1})]),_:1})]),_:1}))}});export{ie as default};
