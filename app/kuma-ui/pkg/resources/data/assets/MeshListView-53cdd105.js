import{g as w,W as z,Q as $,D as S,u as C,J as x}from"./kongponents.es-fed304fd.js";import{d as v,B as A,q as f,s as V,o as h,a as k,l as B,j as L,w as s,e as N,n as g,b as o,r as D,h as i,k as b,g as y,t as M}from"./index-f4ec0be6.js";import{f as R,i as j,A as q,_ as E}from"./RouteView.vue_vue_type_script_setup_true_lang-09fd82a0.js";import{_ as I}from"./DataSource.vue_vue_type_script_setup_true_lang-1e42b670.js";import{_ as J}from"./RouteTitle.vue_vue_type_script_setup_true_lang-d39e0ee5.js";const Q={key:0,class:"app-collection-toolbar"},W=v({__name:"AppCollection",props:{total:{default:0},pageNumber:{},pageSize:{},items:{},error:{}},emits:["change"],setup(c,{emit:n}){const t=c,d=A(),m=f(t.items),l=f(0);V(()=>t.items,()=>{l.value++,m.value=t.items});const p=a=>{const r=a.target.closest("tr");if(r){const e=r.querySelector("a");e!==null&&e.click()}};return(a,r)=>(h(),k(o(w),{class:"app-collection","has-error":typeof t.error<"u","pagination-total-items":t.total,"initial-fetcher-params":{page:t.pageNumber,pageSize:t.pageSize},"fetcher-cache-key":String(l.value),fetcher:({page:e,pageSize:u,query:_})=>(n("change",{page:e,size:u,s:_}),{data:m.value}),"cell-attrs":({headerKey:e})=>({class:`${e}-column`}),"empty-state-icon-size":"96","disable-sorting":"","hide-pagination-when-optional":"","onRow:click":p},B({_:2},[L(Object.keys(o(d)),e=>({name:e,fn:s(({row:u,rowValue:_})=>[e==="toolbar"?(h(),N("div",Q,[g(a.$slots,"toolbar",{},void 0,!0)])):g(a.$slots,e,{key:1,row:u,rowValue:_},void 0,!0)])}))]),1032,["has-error","pagination-total-items","initial-fetcher-params","fetcher-cache-key","fetcher","cell-attrs"]))}});const K=R(W,[["__scopeId","data-v-7e51f7e1"]]),O={class:"stack"},U=v({__name:"MeshListView",props:{page:{},size:{}},setup(c){const n=c,{t}=j();return(d,m)=>{const l=D("RouterLink");return h(),k(E,{name:"mesh-list-view"},{default:s(({route:p})=>[i(I,{src:`/meshes?page=${n.page}&size=${n.size}`},{default:s(({data:a,error:r})=>[i(q,{breadcrumbs:[{to:{name:"mesh-list-view"},text:o(t)("meshes.routes.items.breadcrumbs")}]},{title:s(()=>[b("h1",null,[i(J,{title:o(t)("meshes.routes.items.title"),render:!0},null,8,["title"])])]),default:s(()=>[y(),b("div",O,[i(o(z),null,{body:s(()=>[i(K,{"data-testid":"mesh-collection","empty-state-title":o(t)("common.emptyState.title"),"empty-state-message":o(t)("common.emptyState.message",{type:"Meshes"}),headers:[{label:"Name",key:"name"},{label:"Actions",key:"actions",hideLabel:!0}],"page-number":n.page,"page-size":n.size,total:a==null?void 0:a.total,items:a==null?void 0:a.items,error:r,onChange:p.update},{name:s(({row:e})=>[i(l,{to:{name:"mesh-detail-view",params:{mesh:e.name}}},{default:s(()=>[y(M(e.name),1)]),_:2},1032,["to"])]),actions:s(({row:e})=>[i(o($),{class:"actions-dropdown","kpop-attributes":{placement:"bottomEnd",popoverClasses:"mt-5 more-actions-popover"},width:"150"},{default:s(()=>[i(o(S),{class:"non-visual-button",appearance:"secondary",size:"small"},{icon:s(()=>[i(o(C),{color:"var(--black-400)",icon:"more",size:"16"})]),_:1})]),items:s(()=>[i(o(x),{item:{to:{name:"mesh-detail-view",params:{mesh:e.name}},label:"View"}},null,8,["item"])]),_:2},1024)]),_:2},1032,["empty-state-title","empty-state-message","page-number","page-size","total","items","error","onChange"])]),_:2},1024)])]),_:2},1032,["breadcrumbs"])]),_:2},1032,["src"])]),_:1})}}});export{U as default};
