import{d as P,j as b,c as w,A as S,o as n,a as y,w as a,h as s,b as k,g as u,B as N,q as o,t as c,e as d,$ as T,F}from"./index-840bc7ec.js";import{L as I}from"./LoadingBox-7713ccf4.js";import{O as E,a as $,b as L}from"./OnboardingPage-259f1d88.js";import{m as V,g as C,A as H,S as M,_ as R,f as j}from"./RouteView.vue_vue_type_script_setup_true_lang-3639338c.js";import{_ as q}from"./RouteTitle.vue_vue_type_script_setup_true_lang-6980b8ac.js";import{g as z}from"./dataplane-30467516.js";const K={key:0,class:"status-loading-box mb-4"},U={key:1},W={class:"mb-4"},G=P({__name:"DataplanesOverview",setup(J){const p=V(),{t:x}=C(),D=[{label:"Mesh",key:"mesh"},{label:"Name",key:"name"},{label:"Status",key:"status"}],e=b({total:0,data:[]}),l=b(null),A=w(()=>e.value.data.length>0?"Success":"Waiting for DPPs"),m=w(()=>e.value.data.length>0?"The following data plane proxies (DPPs) are connected to the control plane:":null);S(function(){_()}),f();function _(){l.value!==null&&window.clearTimeout(l.value)}async function f(){let i=!1;const r=[];try{const{items:t}=await p.getAllDataplanes({size:10});if(Array.isArray(t))for(const B of t){const{name:v,mesh:g}=B,O=await p.getDataplaneOverviewFromMesh({mesh:g,name:v}),h=z(O.dataplaneInsight);h==="offline"&&(i=!0),r.push({status:h,name:v,mesh:g})}}catch(t){console.error(t)}e.value.data=r,e.value.total=e.value.data.length,i&&(_(),l.value=window.setTimeout(f,1e3))}return(i,r)=>(n(),y(R,null,{default:a(()=>[s(q,{title:k(x)("onboarding.routes.dataplanes-overview.title")},null,8,["title"]),u(),s(H,null,{default:a(()=>[s(E,null,{header:a(()=>[s($,null,N({title:a(()=>[o("p",null,c(A.value),1)]),_:2},[m.value!==null?{name:"description",fn:a(()=>[o("p",null,c(m.value),1)]),key:"0"}:void 0]),1024)]),content:a(()=>[e.value.data.length===0?(n(),d("div",K,[s(I)])):(n(),d("div",U,[o("p",W,[o("b",null,"Found "+c(e.value.data.length)+" DPPs:",1)]),u(),s(k(T),{class:"mb-4",fetcher:()=>e.value,headers:D,"disable-pagination":""},{status:a(({rowValue:t})=>[t?(n(),y(M,{key:0,status:t},null,8,["status"])):(n(),d(F,{key:1},[u(`
                  —
                `)],64))]),_:1},8,["fetcher"])]))]),navigation:a(()=>[s(L,{"next-step":"onboarding-completed","previous-step":"onboarding-add-services-code","should-allow-next":e.value.data.length>0},null,8,["should-allow-next"])]),_:1})]),_:1})]),_:1}))}});const te=j(G,[["__scopeId","data-v-4588fbe4"]]);export{te as default};
