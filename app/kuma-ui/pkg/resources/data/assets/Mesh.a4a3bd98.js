import{_ as A,M as V,k as L,cM as C,K as B,o as c,f as N,h as t,d as m,w as l,t as S,e as M,i as d,v as h,b as g,n as u,cN as k,cO as U,r as _}from"./index.6180ff6f.js";import{F as D}from"./FormatForCLI.231d97e4.js";import{F as x,S as I,E as O}from"./EntityScanner.8a022c02.js";import{T as R}from"./TabsWidget.6e628a3c.js";import{C as F}from"./CodeView.4223104c.js";import"./index.124532c3.js";import"./datadogEvents.f0c2b8e0.js";import"./CodeBlock.c76d2442.js";function P(o,n){return Object.keys(o).filter(s=>!n.includes(s)).map(s=>Object.assign({},{[s]:o[s]})).reduce((s,v)=>Object.assign(s,v),{})}const K={mtls:{enabledBackend:null,backends:[]},tracing:{defaultBackend:null,backends:[{name:null,type:null}]},logging:{backends:[{name:null,format:'{ "destination": "%KUMA_DESTINATION_SERVICE%", "destinationAddress": "%UPSTREAM_LOCAL_ADDRESS%", "source": "%KUMA_SOURCE_SERVICE%", "sourceAddress": "%KUMA_SOURCE_ADDRESS%", "bytesReceived": "%BYTES_RECEIVED%", "bytesSent": "%BYTES_SENT%"}',type:null}]},metrics:{enabledBackend:null,backends:[{name:null,type:null}]}},z={name:"MeshWizard",components:{FormFragment:x,TabsWidget:R,StepSkeleton:I,CodeView:F,EntityScanner:O},mixins:[D],data(){return{productName:V,selectedTab:"",schema:K,steps:[{label:"General & Security",slug:"general"},{label:"Logging",slug:"logging"},{label:"Tracing",slug:"tracing"},{label:"Metrics",slug:"metrics"},{label:"Install",slug:"complete"}],tabs:[{hash:"#kubernetes",title:"Kubernetes"},{hash:"#universal",title:"Universal"}],sidebarContent:[{name:"mesh"},{name:"did-you-know"}],formConditions:{mtlsEnabled:!1,loggingEnabled:!1,tracingEnabled:!1,metricsEnabled:!1,loggingType:null},startScanner:!1,scanFound:!1,hideScannerSiblings:!1,scanError:!1,isComplete:!1,validate:{meshName:"",meshCAName:"",meshLoggingBackend:"",meshTracingBackend:"",meshMetricsName:"",meshTracingZipkinURL:"",mtlsEnabled:"disabled",meshCA:"builtin",loggingEnabled:"disabled",loggingType:"tcp",meshLoggingPath:"/",meshLoggingAddress:"127.0.0.1:5000",meshLoggingBackendFormat:"{ start_time: '%START_TIME%', source: '%KUMA_SOURCE_SERVICE%', destination: '%KUMA_DESTINATION_SERVICE%', source_address: '%KUMA_SOURCE_ADDRESS_WITHOUT_PORT%', destination_address: '%UPSTREAM_HOST%', duration_millis: '%DURATION%', bytes_received: '%BYTES_RECEIVED%', bytes_sent: '%BYTES_SENT%' }",tracingEnabled:"disabled",meshTracingType:"zipkin",meshTracingSampling:99.9,metricsEnabled:"disabled",meshMetricsType:"prometheus",meshMetricsDataplanePort:5670,meshMetricsDataplanePath:"/metrics"},vmsg:[],utm:"?utm_source=Kuma&utm_medium=Kuma-GUI"}},computed:{...L({title:"config/getTagline",kumaDocsVersion:"config/getKumaDocsVersion",environment:"config/getEnvironment"}),codeOutput(){const o=this.schema,n=Object.assign({},o),s=this.validate;if(!s)return;const v=s.mtlsEnabled==="enabled",e=s.loggingEnabled==="enabled",p=s.tracingEnabled==="enabled",E=s.metricsEnabled==="enabled",r={mtls:v,logging:e,tracing:p,metrics:E},b=[];if(Object.entries(r).forEach(i=>{const f=i[1],y=i[0];f?b.filter(a=>a!==y):b.push(y)}),v){n.mtls.enabled=!0;const i=n.mtls,f=this.validate.meshCA,y=this.validate.meshCAName;i.backends=[],i.enabledBackend=y,f==="provided"?i.backends=[{name:y,type:f,conf:{cert:{secret:""},key:{secret:""}}}]:i.backends=[{name:y,type:f}]}if(e){const i=n.logging.backends[0],f=i.format;i.conf={},i.name=s.meshLoggingBackend,i.type=s.loggingType,i.format=s.meshLoggingBackendFormat||f,s.loggingType==="tcp"?i.conf.address=s.meshLoggingAddress||"127.0.0.1:5000":s.loggingType==="file"&&(i.conf.path=s.meshLoggingPath)}if(p){const i=n.tracing;i.defaultBackend=s.meshTracingBackend,i.backends[0].type=s.meshTracingType||"zipkin",i.backends[0].name=s.meshTracingBackend,i.backends[0].sampling=s.meshTracingSampling||100,i.backends[0].conf={},i.backends[0].conf.url=s.meshTracingZipkinURL}if(E){const i=n.metrics;i.backends[0].conf={},i.enabledBackend=s.meshMetricsName,i.backends[0].type=s.meshMetricsType||"prometheus",i.backends[0].name=s.meshMetricsName,i.backends[0].conf.port=s.meshMetricsDataplanePort||5670,i.backends[0].conf.path=s.meshMetricsDataplanePath||"/metrics"}const T=P(n,b);let w;return this.selectedTab==="#kubernetes"?w={apiVersion:"kuma.io/v1alpha1",kind:"Mesh",metadata:{name:s.meshName},spec:T}:w={type:"Mesh",name:s.meshName,...T},this.formatForCLI(w,'" | kumactl apply -f -')},nextDisabled(){const{meshName:o,meshCAName:n,meshLoggingBackend:s,meshTracingBackend:v,meshTracingZipkinURL:e,meshMetricsName:p,mtlsEnabled:E,loggingEnabled:r,tracingEnabled:b,metricsEnabled:T,meshLoggingPath:w,loggingType:i}=this.validate;return!o.length||E==="enabled"&&!n?!0:this.$route.query.step==="1"?r==="disabled"?!1:s?i==="file"&&!w:!0:this.$route.query.step==="2"?b==="enabled"&&!(v&&e):this.$route.query.step==="3"?T==="enabled"&&!p:!1}},watch:{"validate.meshName"(o){const n=C(o);this.validate.meshName=n,this.validateMeshName(n)},"validate.meshCAName"(o){this.validate.meshCAName=C(o)},"validate.meshLoggingBackend"(o){this.validate.meshLoggingBackend=C(o)},"validate.meshTracingBackend"(o){this.validate.meshTracingBackend=C(o)},"validate.meshMetricsName"(o){this.validate.meshMetricsName=C(o)}},methods:{onTabChange(o){this.selectedTab=o},hideSiblings(){this.hideScannerSiblings=!0},validateMeshName(o){!o||o===""?this.vmsg.meshName="A Mesh name is required to proceed":this.vmsg.meshName=""},scanForEntity(){const o=this.validate.meshName;this.scanComplete=!1,this.scanError=!1,o&&B.getMesh({name:o}).then(n=>{n&&n.name.length>0?(this.isRunning=!0,this.scanFound=!0):this.scanError=!0}).catch(n=>{this.scanError=!0,console.error(n)}).finally(()=>{this.scanComplete=!0})}}},Y={class:"wizard"},j={class:"wizard__content"},W=t("code",null,"kubectl",-1),G=M(" (if you are running in Kubernetes mode) or "),Z=t("code",null,"kumactl",-1),q=M(" / API (if you are running in Universal mode). "),H=t("h3",null," To get started, please fill in the following information: ",-1),J={class:"k-input-label mx-2"},Q=t("span",null,"Disabled",-1),X={class:"k-input-label mx-2"},$=t("span",null,"Enabled",-1),ee=t("option",{value:"builtin"}," builtin ",-1),te=t("option",{value:"provided"}," provided ",-1),ne=[ee,te],ae=t("p",{class:"help"}," If you've enabled mTLS, you must select a CA. ",-1),le=t("h3",null," Setup Logging ",-1),se=t("p",null,' You can setup as many logging backends as you need that you can later use to log traffic via the "TrafficLog" policy. In this wizard, we allow you to configure one backend, but you can add more manually if you wish. ',-1),ie={class:"k-input-label mx-2"},oe=t("span",null,"Disabled",-1),re={class:"k-input-label mx-2"},de=t("span",null,"Enabled",-1),ce={key:1},me=t("option",{value:"tcp"}," TCP ",-1),ue=t("option",{value:"file"}," File ",-1),ge=[me,ue],pe=t("h3",null," Setup Tracing ",-1),he=t("p",null,' You can setup as many tracing backends as you need that you can later use to log traffic via the "TrafficTrace" policy. In this wizard we allow you to configure one backend, but you can add more manually as you wish. ',-1),be={class:"k-input-label mx-2"},fe=t("span",null,"Disabled",-1),ye={class:"k-input-label mx-2"},ke=t("span",null,"Enabled",-1),_e=t("option",{value:"zipkin"}," Zipkin ",-1),ve=[_e],Ee=t("h3",null," Setup Metrics ",-1),Te=t("p",null," You can expose metrics from every data-plane on a configurable path and port that a metrics service, like Prometheus, can use to fetch them. ",-1),we={class:"k-input-label mx-2"},Se=t("span",null,"Disabled",-1),Me={class:"k-input-label mx-2"},Ce=t("span",null,"Enabled",-1),Ne=t("option",{value:"prometheus"}," Prometheus ",-1),Ue=[Ne],Ae={key:0},Ve={key:0},Le=t("h3",null," Install a new Mesh ",-1),Be=t("h3",null,"Searching\u2026",-1),De=t("p",null,"We are looking for your mesh.",-1),xe=t("h3",null,"Done!",-1),Ie=M(" Your Mesh "),Oe={key:0},Re=M(" was found! "),Fe=M(" See Meshes "),Pe=t("h3",null,"Mesh not found",-1),Ke=t("p",null,"We were unable to find your mesh.",-1),ze=t("p",null," You haven't filled any data out yet! Please return to the first step and fill out your information. ",-1),Ye=t("h3",null,"Mesh",-1),je=["href"],We=t("h3",null,"Did You Know?",-1),Ge=t("p",null," As you know, the GUI is read-only, but it will be providing instructions to create a new Mesh and verify everything worked well. ",-1);function Ze(o,n,s,v,e,p){const E=_("KAlert"),r=_("FormFragment"),b=_("KCard"),T=_("CodeView"),w=_("TabsWidget"),i=_("KButton"),f=_("EntityScanner"),y=_("StepSkeleton");return c(),N("div",Y,[t("div",j,[m(y,{steps:e.steps,"sidebar-content":e.sidebarContent,"footer-enabled":e.hideScannerSiblings===!1,"next-disabled":p.nextDisabled},{general:l(()=>[t("p",null," Welcome to the wizard for creating a new Mesh resource in "+S(e.productName)+". We will be providing you with a few steps that will get you started. ",1),t("p",null,[M(" As you know, the "+S(e.productName)+" GUI is read-only, so at the end of this wizard we will be generating the configuration that you can apply with either ",1),W,G,Z,q]),H,m(b,{class:"my-6 k-card--small",title:"Mesh Information","has-shadow":""},{body:l(()=>[m(r,{title:"Mesh name","for-attr":"mesh-name"},{default:l(()=>[d(t("input",{id:"mesh-name","onUpdate:modelValue":n[0]||(n[0]=a=>e.validate.meshName=a),type:"text",class:"k-input w-100",placeholder:"your-mesh-name",required:""},null,512),[[h,e.validate.meshName]]),e.vmsg.meshName?(c(),g(E,{key:0,appearance:"danger",size:"small","alert-message":e.vmsg.meshName},null,8,["alert-message"])):u("",!0)]),_:1}),m(r,{title:"Mutual TLS"},{default:l(()=>[t("label",J,[d(t("input",{ref:"mtlsDisabled","onUpdate:modelValue":n[1]||(n[1]=a=>e.validate.mtlsEnabled=a),value:"disabled",name:"mtls",type:"radio",class:"k-input mr-2"},null,512),[[k,e.validate.mtlsEnabled]]),Q]),t("label",X,[d(t("input",{id:"mtls-enabled","onUpdate:modelValue":n[2]||(n[2]=a=>e.validate.mtlsEnabled=a),value:"enabled",name:"mtls",type:"radio",class:"k-input mr-2"},null,512),[[k,e.validate.mtlsEnabled]]),$])]),_:1}),e.validate.mtlsEnabled==="enabled"?(c(),g(r,{key:0,title:"Certificate name","for-attr":"certificate-name"},{default:l(()=>[d(t("input",{id:"certificate-name","onUpdate:modelValue":n[3]||(n[3]=a=>e.validate.meshCAName=a),type:"text",class:"k-input w-100",placeholder:"your-certificate-name"},null,512),[[h,e.validate.meshCAName]])]),_:1})):u("",!0),e.validate.mtlsEnabled==="enabled"?(c(),g(r,{key:1,title:"Certificate Authority","for-attr":"certificate-authority"},{default:l(()=>[d(t("select",{id:"certificate-authority","onUpdate:modelValue":n[4]||(n[4]=a=>e.validate.meshCA=a),class:"k-input w-100",name:"certificate-authority"},ne,512),[[U,e.validate.meshCA]]),ae]),_:1})):u("",!0)]),_:1})]),logging:l(()=>[le,se,m(b,{class:"my-6 k-card--small",title:"Logging Configuration","has-shadow":""},{body:l(()=>[m(r,{title:"Logging"},{default:l(()=>[t("label",ie,[d(t("input",{id:"logging-disabled","onUpdate:modelValue":n[5]||(n[5]=a=>e.validate.loggingEnabled=a),value:"disabled",name:"logging",type:"radio",class:"k-input mr-2"},null,512),[[k,e.validate.loggingEnabled]]),oe]),t("label",re,[d(t("input",{id:"logging-enabled","onUpdate:modelValue":n[6]||(n[6]=a=>e.validate.loggingEnabled=a),value:"enabled",name:"logging",type:"radio",class:"k-input mr-2"},null,512),[[k,e.validate.loggingEnabled]]),de])]),_:1}),e.validate.loggingEnabled==="enabled"?(c(),g(r,{key:0,title:"Backend name","for-attr":"backend-name"},{default:l(()=>[d(t("input",{id:"backend-name","onUpdate:modelValue":n[7]||(n[7]=a=>e.validate.meshLoggingBackend=a),type:"text",class:"k-input w-100",placeholder:"your-backend-name"},null,512),[[h,e.validate.meshLoggingBackend]])]),_:1})):u("",!0),e.validate.loggingEnabled==="enabled"?(c(),N("div",ce,[m(r,{title:"Type"},{default:l(()=>[d(t("select",{id:"logging-type",ref:"loggingTypeSelect","onUpdate:modelValue":n[8]||(n[8]=a=>e.validate.loggingType=a),class:"k-input w-100",name:"logging-type"},ge,512),[[U,e.validate.loggingType]])]),_:1}),e.validate.loggingType==="file"?(c(),g(r,{key:0,title:"Path","for-attr":"backend-address"},{default:l(()=>[d(t("input",{id:"backend-address","onUpdate:modelValue":n[9]||(n[9]=a=>e.validate.meshLoggingPath=a),type:"text",class:"k-input w-100"},null,512),[[h,e.validate.meshLoggingPath]])]),_:1})):u("",!0),e.validate.loggingType==="tcp"?(c(),g(r,{key:1,title:"Address","for-attr":"backend-address"},{default:l(()=>[d(t("input",{id:"backend-address","onUpdate:modelValue":n[10]||(n[10]=a=>e.validate.meshLoggingAddress=a),type:"text",class:"k-input w-100"},null,512),[[h,e.validate.meshLoggingAddress]])]),_:1})):u("",!0),m(r,{title:"Format","for-attr":"backend-format"},{default:l(()=>[d(t("textarea",{id:"backend-format","onUpdate:modelValue":n[11]||(n[11]=a=>e.validate.meshLoggingBackendFormat=a),class:"k-input w-100 code-sample",rows:"12"},null,512),[[h,e.validate.meshLoggingBackendFormat]])]),_:1})])):u("",!0)]),_:1})]),tracing:l(()=>[pe,he,m(b,{class:"my-6 k-card--small",title:"Tracing Configuration","has-shadow":""},{body:l(()=>[m(r,{title:"Tracing"},{default:l(()=>[t("label",be,[d(t("input",{id:"tracing-disabled","onUpdate:modelValue":n[12]||(n[12]=a=>e.validate.tracingEnabled=a),value:"disabled",name:"tracing",type:"radio",class:"k-input mr-2"},null,512),[[k,e.validate.tracingEnabled]]),fe]),t("label",ye,[d(t("input",{id:"tracing-enabled","onUpdate:modelValue":n[13]||(n[13]=a=>e.validate.tracingEnabled=a),value:"enabled",name:"tracing",type:"radio",class:"k-input mr-2"},null,512),[[k,e.validate.tracingEnabled]]),ke])]),_:1}),e.validate.tracingEnabled==="enabled"?(c(),g(r,{key:0,title:"Backend name","for-attr":"tracing-backend-name"},{default:l(()=>[d(t("input",{id:"tracing-backend-name","onUpdate:modelValue":n[14]||(n[14]=a=>e.validate.meshTracingBackend=a),type:"text",class:"k-input w-100",placeholder:"your-tracing-backend-name"},null,512),[[h,e.validate.meshTracingBackend]])]),_:1})):u("",!0),e.validate.tracingEnabled==="enabled"?(c(),g(r,{key:1,title:"Type","for-attr":"tracing-type"},{default:l(()=>[d(t("select",{id:"tracing-type","onUpdate:modelValue":n[15]||(n[15]=a=>e.validate.meshTracingType=a),class:"k-input w-100",name:"tracing-type"},ve,512),[[U,e.validate.meshTracingType]])]),_:1})):u("",!0),e.validate.tracingEnabled==="enabled"?(c(),g(r,{key:2,title:"Sampling","for-attr":"tracing-sampling"},{default:l(()=>[d(t("input",{id:"tracing-sampling","onUpdate:modelValue":n[16]||(n[16]=a=>e.validate.meshTracingSampling=a),type:"number",class:"k-input w-100",step:"0.1",min:"0",max:"100"},null,512),[[h,e.validate.meshTracingSampling]])]),_:1})):u("",!0),e.validate.tracingEnabled==="enabled"?(c(),g(r,{key:3,title:"URL","for-attr":"tracing-zipkin-url"},{default:l(()=>[d(t("input",{id:"tracing-zipkin-url","onUpdate:modelValue":n[17]||(n[17]=a=>e.validate.meshTracingZipkinURL=a),type:"text",class:"k-input w-100",placeholder:"http://zipkin.url:1234"},null,512),[[h,e.validate.meshTracingZipkinURL]])]),_:1})):u("",!0)]),_:1})]),metrics:l(()=>[Ee,Te,m(b,{class:"my-6 k-card--small",title:"Metrics Configuration","has-shadow":""},{body:l(()=>[m(r,{title:"Metrics"},{default:l(()=>[t("label",we,[d(t("input",{id:"metrics-disabled","onUpdate:modelValue":n[18]||(n[18]=a=>e.validate.metricsEnabled=a),value:"disabled",name:"metrics",type:"radio",class:"k-input mr-2"},null,512),[[k,e.validate.metricsEnabled]]),Se]),t("label",Me,[d(t("input",{id:"metrics-enabled","onUpdate:modelValue":n[19]||(n[19]=a=>e.validate.metricsEnabled=a),value:"enabled",name:"metrics",type:"radio",class:"k-input mr-2"},null,512),[[k,e.validate.metricsEnabled]]),Ce])]),_:1}),e.validate.metricsEnabled==="enabled"?(c(),g(r,{key:0,title:"Backend name","for-attr":"metrics-name"},{default:l(()=>[d(t("input",{id:"metrics-name","onUpdate:modelValue":n[20]||(n[20]=a=>e.validate.meshMetricsName=a),type:"text",class:"k-input w-100",placeholder:"your-metrics-backend-name"},null,512),[[h,e.validate.meshMetricsName]])]),_:1})):u("",!0),e.validate.metricsEnabled==="enabled"?(c(),g(r,{key:1,title:"Type","for-attr":"metrics-type"},{default:l(()=>[d(t("select",{id:"metrics-type","onUpdate:modelValue":n[21]||(n[21]=a=>e.validate.meshMetricsType=a),class:"k-input w-100",name:"metrics-type"},Ue,512),[[U,e.validate.meshMetricsType]])]),_:1})):u("",!0),e.validate.metricsEnabled==="enabled"?(c(),g(r,{key:2,title:"Dataplane port","for-attr":"metrics-dataplane-port"},{default:l(()=>[d(t("input",{id:"metrics-dataplane-port","onUpdate:modelValue":n[22]||(n[22]=a=>e.validate.meshMetricsDataplanePort=a),type:"number",class:"k-input w-100",step:"1",min:"0",max:"65535",placeholder:"1234"},null,512),[[h,e.validate.meshMetricsDataplanePort]])]),_:1})):u("",!0),e.validate.metricsEnabled==="enabled"?(c(),g(r,{key:3,title:"Dataplane path","for-attr":"metrics-dataplane-path"},{default:l(()=>[d(t("input",{id:"metrics-dataplane-path","onUpdate:modelValue":n[23]||(n[23]=a=>e.validate.meshMetricsDataplanePath=a),type:"text",class:"k-input w-100"},null,512),[[h,e.validate.meshMetricsDataplanePath]])]),_:1})):u("",!0)]),_:1})]),complete:l(()=>[p.codeOutput?(c(),N("div",Ae,[e.hideScannerSiblings===!1?(c(),N("div",Ve,[Le,t("p",null," Since the "+S(e.productName)+" GUI is read-only mode to follow Ops best practices, please execute the following command in your shell to create the entity. "+S(e.productName)+" will automatically detect when the new entity has been created. ",1),m(w,{loaders:!1,tabs:e.tabs,"initial-tab-override":o.environment,onOnTabChange:p.onTabChange},{kubernetes:l(()=>[m(T,{title:"Kubernetes","copy-button-text":"Copy Command to Clipboard",lang:"bash",content:p.codeOutput},null,8,["content"])]),universal:l(()=>[m(T,{title:"Universal","copy-button-text":"Copy Command to Clipboard",lang:"bash",content:p.codeOutput},null,8,["content"])]),_:1},8,["tabs","initial-tab-override","onOnTabChange"])])):u("",!0),m(f,{"loader-function":p.scanForEntity,"should-start":!0,"has-error":e.scanError,"can-complete":e.scanFound,onHideSiblings:p.hideSiblings},{"loading-title":l(()=>[Be]),"loading-content":l(()=>[De]),"complete-title":l(()=>[xe]),"complete-content":l(()=>[t("p",null,[Ie,e.validate.meshName?(c(),N("strong",Oe,S(e.validate.meshName),1)):u("",!0),Re]),t("p",null,[m(i,{appearance:"primary",to:{name:"all-meshes"}},{default:l(()=>[Fe]),_:1})])]),"error-title":l(()=>[Pe]),"error-content":l(()=>[Ke]),_:1},8,["loader-function","has-error","can-complete","onHideSiblings"])])):(c(),g(E,{key:1,appearance:"danger"},{alertMessage:l(()=>[ze]),_:1}))]),mesh:l(()=>[Ye,t("p",null," In "+S(o.title)+", a Mesh resource allows you to define an isolated environment for your data-planes and policies. It's isolated because the mTLS CA you choose can be different from the one configured for our Meshes. Ideally, you will have either a large Mesh with all the workloads, or one Mesh per application for better isolation. ",1),t("p",null,[t("a",{href:`https://kuma.io/docs/${o.kumaDocsVersion}/policies/mesh/${e.utm}`,target:"_blank"}," Learn More ",8,je)])]),"did-you-know":l(()=>[We,Ge]),_:1},8,["steps","sidebar-content","footer-enabled","next-disabled"])])])}const nt=A(z,[["render",Ze]]);export{nt as default};
