import{d as p,a2 as f,o as e,j as i,g as s,w as n,m as _,D as m,i as t,ab as k,E as g,t as v,ae as q,F as E,h as o,W as h,aq as x,K as S,l as u,p as b,q as B}from"./index-35b33747.js";import{_ as C}from"./CodeBlock.vue_vue_type_style_index_0_lang-cbb8f36b.js";const D={class:"envoy-data-actions"},I=p({__name:"EnvoyData",props:{status:{type:String,required:!0},resource:{type:String,required:!0},src:{type:String,required:!0},queryKey:{type:String,required:!0}},setup(l){const r=l,{t:y}=f();return(N,K)=>(e(),i("div",null,[r.status!=="online"?(e(),s(t(k),{key:0,appearance:"info"},{alertMessage:n(()=>[_("p",null,m(t(y)("common.detail.no_envoy_data",{resource:r.resource})),1)]),_:1})):(e(),s(b,{key:1,src:r.src},{default:n(({data:a,error:c,refresh:d})=>[c?(e(),s(g,{key:0,error:c},null,8,["error"])):a===void 0?(e(),s(v,{key:1})):a===""?(e(),s(q,{key:2})):(e(),i(E,{key:3},[_("div",D,[o(t(h),{appearance:"primary","data-testid":"envoy-data-refresh-button",onClick:d},{default:n(()=>[o(t(x),{size:t(S)},null,8,["size"]),u(`

            Refresh
          `)]),_:2},1032,["onClick"])]),u(),o(C,{id:"code-block-envoy-data",language:"json",code:typeof a=="string"?a:JSON.stringify(a,null,2),"is-searchable":"","query-key":r.queryKey},null,8,["code","query-key"])],64))]),_:1},8,["src"]))]))}});const j=B(I,[["__scopeId","data-v-faac85b9"]]);export{j as E};
