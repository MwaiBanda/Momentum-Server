/**
 * @remix-run/router v1.8.0
 *
 * Copyright (c) Remix Software Inc.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE.md file in the root directory of this source tree.
 *
 * @license MIT
 */function T(){return T=Object.assign?Object.assign.bind():function(e){for(var t=1;t<arguments.length;t++){var r=arguments[t];for(var n in r)Object.prototype.hasOwnProperty.call(r,n)&&(e[n]=r[n])}return e},T.apply(this,arguments)}var H;(function(e){e.Pop="POP",e.Push="PUSH",e.Replace="REPLACE"})(H||(H={}));const ct="popstate";function zr(e){e===void 0&&(e={});function t(n,a){let{pathname:s,search:d,hash:u}=n.location;return Fe("",{pathname:s,search:d,hash:u},a.state&&a.state.usr||null,a.state&&a.state.key||"default")}function r(n,a){return typeof a=="string"?a:Ae(a)}return Yt(t,r,null,e)}function U(e,t){if(e===!1||e===null||typeof e>"u")throw new Error(t)}function ge(e,t){if(!e){typeof console<"u"&&console.warn(t);try{throw new Error(t)}catch{}}}function Vt(){return Math.random().toString(36).substr(2,8)}function ut(e,t){return{usr:e.state,key:e.key,idx:t}}function Fe(e,t,r,n){return r===void 0&&(r=null),T({pathname:typeof e=="string"?e:e.pathname,search:"",hash:""},typeof t=="string"?se(t):t,{state:r,key:t&&t.key||n||Vt()})}function Ae(e){let{pathname:t="/",search:r="",hash:n=""}=e;return r&&r!=="?"&&(t+=r.charAt(0)==="?"?r:"?"+r),n&&n!=="#"&&(t+=n.charAt(0)==="#"?n:"#"+n),t}function se(e){let t={};if(e){let r=e.indexOf("#");r>=0&&(t.hash=e.substr(r),e=e.substr(0,r));let n=e.indexOf("?");n>=0&&(t.search=e.substr(n),e=e.substr(0,n)),e&&(t.pathname=e)}return t}function Yt(e,t,r,n){n===void 0&&(n={});let{window:a=document.defaultView,v5Compat:s=!1}=n,d=a.history,u=H.Pop,f=null,h=m();h==null&&(h=0,d.replaceState(T({},d.state,{idx:h}),""));function m(){return(d.state||{idx:null}).idx}function M(){u=H.Pop;let v=m(),S=v==null?null:v-h;h=v,f&&f({action:u,location:x.location,delta:S})}function D(v,S){u=H.Push;let B=Fe(x.location,v,S);r&&r(B,v),h=m()+1;let K=ut(B,h),l=x.createHref(B);try{d.pushState(K,"",l)}catch(L){if(L instanceof DOMException&&L.name==="DataCloneError")throw L;a.location.assign(l)}s&&f&&f({action:u,location:x.location,delta:1})}function C(v,S){u=H.Replace;let B=Fe(x.location,v,S);r&&r(B,v),h=m();let K=ut(B,h),l=x.createHref(B);d.replaceState(K,"",l),s&&f&&f({action:u,location:x.location,delta:0})}function E(v){let S=a.location.origin!=="null"?a.location.origin:a.location.href,B=typeof v=="string"?v:Ae(v);return U(S,"No window.location.(origin|href) available to create URL for href: "+B),new URL(B,S)}let x={get action(){return u},get location(){return e(a,d)},listen(v){if(f)throw new Error("A history only accepts one active listener");return a.addEventListener(ct,M),f=v,()=>{a.removeEventListener(ct,M),f=null}},createHref(v){return t(a,v)},createURL:E,encodeLocation(v){let S=E(v);return{pathname:S.pathname,search:S.search,hash:S.hash}},push:D,replace:C,go(v){return d.go(v)}};return x}var z;(function(e){e.data="data",e.deferred="deferred",e.redirect="redirect",e.error="error"})(z||(z={}));const Gt=new Set(["lazy","caseSensitive","path","id","index","children"]);function Jt(e){return e.index===!0}function Ye(e,t,r,n){return r===void 0&&(r=[]),n===void 0&&(n={}),e.map((a,s)=>{let d=[...r,s],u=typeof a.id=="string"?a.id:d.join("-");if(U(a.index!==!0||!a.children,"Cannot specify children on an index route"),U(!n[u],'Found a route id collision on id "'+u+`".  Route id's must be globally unique within Data Router usages`),Jt(a)){let f=T({},a,t(a),{id:u});return n[u]=f,f}else{let f=T({},a,t(a),{id:u,children:void 0});return n[u]=f,a.children&&(f.children=Ye(a.children,t,d,n)),f}})}function Se(e,t,r){r===void 0&&(r="/");let n=typeof t=="string"?se(t):t,a=ze(n.pathname||"/",r);if(a==null)return null;let s=Mt(e);Xt(s);let d=null;for(let u=0;d==null&&u<s.length;++u)d=ir(s[u],sr(a));return d}function Mt(e,t,r,n){t===void 0&&(t=[]),r===void 0&&(r=[]),n===void 0&&(n="");let a=(s,d,u)=>{let f={relativePath:u===void 0?s.path||"":u,caseSensitive:s.caseSensitive===!0,childrenIndex:d,route:s};f.relativePath.startsWith("/")&&(U(f.relativePath.startsWith(n),'Absolute route path "'+f.relativePath+'" nested under path '+('"'+n+'" is not valid. An absolute child route path ')+"must start with the combined path of all its parent routes."),f.relativePath=f.relativePath.slice(n.length));let h=Le([n,f.relativePath]),m=r.concat(f);s.children&&s.children.length>0&&(U(s.index!==!0,"Index routes must not have child routes. Please remove "+('all child routes from route path "'+h+'".')),Mt(s.children,t,m,h)),!(s.path==null&&!s.index)&&t.push({path:h,score:ar(h,s.index),routesMeta:m})};return e.forEach((s,d)=>{var u;if(s.path===""||!((u=s.path)!=null&&u.includes("?")))a(s,d);else for(let f of Et(s.path))a(s,d,f)}),t}function Et(e){let t=e.split("/");if(t.length===0)return[];let[r,...n]=t,a=r.endsWith("?"),s=r.replace(/\?$/,"");if(n.length===0)return a?[s,""]:[s];let d=Et(n.join("/")),u=[];return u.push(...d.map(f=>f===""?s:[s,f].join("/"))),a&&u.push(...d),u.map(f=>e.startsWith("/")&&f===""?"/":f)}function Xt(e){e.sort((t,r)=>t.score!==r.score?r.score-t.score:or(t.routesMeta.map(n=>n.childrenIndex),r.routesMeta.map(n=>n.childrenIndex)))}const Qt=/^:\w+$/,Zt=3,er=2,tr=1,rr=10,nr=-2,ft=e=>e==="*";function ar(e,t){let r=e.split("/"),n=r.length;return r.some(ft)&&(n+=nr),t&&(n+=er),r.filter(a=>!ft(a)).reduce((a,s)=>a+(Qt.test(s)?Zt:s===""?tr:rr),n)}function or(e,t){return e.length===t.length&&e.slice(0,-1).every((n,a)=>n===t[a])?e[e.length-1]-t[t.length-1]:0}function ir(e,t){let{routesMeta:r}=e,n={},a="/",s=[];for(let d=0;d<r.length;++d){let u=r[d],f=d===r.length-1,h=a==="/"?t:t.slice(a.length)||"/",m=lr({path:u.relativePath,caseSensitive:u.caseSensitive,end:f},h);if(!m)return null;Object.assign(n,m.params);let M=u.route;s.push({params:n,pathname:Le([a,m.pathname]),pathnameBase:mr(Le([a,m.pathnameBase])),route:M}),m.pathnameBase!=="/"&&(a=Le([a,m.pathnameBase]))}return s}function lr(e,t){typeof e=="string"&&(e={path:e,caseSensitive:!1,end:!0});let[r,n]=dr(e.path,e.caseSensitive,e.end),a=t.match(r);if(!a)return null;let s=a[0],d=s.replace(/(.)\/+$/,"$1"),u=a.slice(1);return{params:n.reduce((h,m,M)=>{if(m==="*"){let D=u[M]||"";d=s.slice(0,s.length-D.length).replace(/(.)\/+$/,"$1")}return h[m]=cr(u[M]||"",m),h},{}),pathname:s,pathnameBase:d,pattern:e}}function dr(e,t,r){t===void 0&&(t=!1),r===void 0&&(r=!0),ge(e==="*"||!e.endsWith("*")||e.endsWith("/*"),'Route path "'+e+'" will be treated as if it were '+('"'+e.replace(/\*$/,"/*")+'" because the `*` character must ')+"always follow a `/` in the pattern. To get rid of this warning, "+('please change the route path to "'+e.replace(/\*$/,"/*")+'".'));let n=[],a="^"+e.replace(/\/*\*?$/,"").replace(/^\/*/,"/").replace(/[\\.*+^$?{}|()[\]]/g,"\\$&").replace(/\/:(\w+)/g,(d,u)=>(n.push(u),"/([^\\/]+)"));return e.endsWith("*")?(n.push("*"),a+=e==="*"||e==="/*"?"(.*)$":"(?:\\/(.+)|\\/*)$"):r?a+="\\/*$":e!==""&&e!=="/"&&(a+="(?:(?=\\/|$))"),[new RegExp(a,t?void 0:"i"),n]}function sr(e){try{return decodeURI(e)}catch(t){return ge(!1,'The URL path "'+e+'" could not be decoded because it is is a malformed URL segment. This is probably due to a bad percent '+("encoding ("+t+").")),e}}function cr(e,t){try{return decodeURIComponent(e)}catch(r){return ge(!1,'The value for the URL param "'+t+'" will not be decoded because'+(' the string "'+e+'" is a malformed URL segment. This is probably')+(" due to a bad percent encoding ("+r+").")),e}}function ze(e,t){if(t==="/")return e;if(!e.toLowerCase().startsWith(t.toLowerCase()))return null;let r=t.endsWith("/")?t.length-1:t.length,n=e.charAt(r);return n&&n!=="/"?null:e.slice(r)||"/"}function ur(e,t){t===void 0&&(t="/");let{pathname:r,search:n="",hash:a=""}=typeof e=="string"?se(e):e;return{pathname:r?r.startsWith("/")?r:fr(r,t):t,search:pr(n),hash:gr(a)}}function fr(e,t){let r=t.replace(/\/+$/,"").split("/");return e.split("/").forEach(a=>{a===".."?r.length>1&&r.pop():a!=="."&&r.push(a)}),r.length>1?r.join("/"):"/"}function qe(e,t,r,n){return"Cannot include a '"+e+"' character in a manually specified "+("`to."+t+"` field ["+JSON.stringify(n)+"].  Please separate it out to the ")+("`to."+r+"` field. Alternatively you may provide the full path as ")+'a string in <Link to="..."> and the router will parse it for you.'}function St(e){return e.filter((t,r)=>r===0||t.route.path&&t.route.path.length>0)}function hr(e,t,r,n){n===void 0&&(n=!1);let a;typeof e=="string"?a=se(e):(a=T({},e),U(!a.pathname||!a.pathname.includes("?"),qe("?","pathname","search",a)),U(!a.pathname||!a.pathname.includes("#"),qe("#","pathname","hash",a)),U(!a.search||!a.search.includes("#"),qe("#","search","hash",a)));let s=e===""||a.pathname==="",d=s?"/":a.pathname,u;if(n||d==null)u=r;else{let M=t.length-1;if(d.startsWith("..")){let D=d.split("/");for(;D[0]==="..";)D.shift(),M-=1;a.pathname=D.join("/")}u=M>=0?t[M]:"/"}let f=ur(a,u),h=d&&d!=="/"&&d.endsWith("/"),m=(s||d===".")&&r.endsWith("/");return!f.pathname.endsWith("/")&&(h||m)&&(f.pathname+="/"),f}const Le=e=>e.join("/").replace(/\/\/+/g,"/"),mr=e=>e.replace(/\/+$/,"").replace(/^\/*/,"/"),pr=e=>!e||e==="?"?"":e.startsWith("?")?e:"?"+e,gr=e=>!e||e==="#"?"":e.startsWith("#")?e:"#"+e;class Lt{constructor(t,r,n,a){a===void 0&&(a=!1),this.status=t,this.statusText=r||"",this.internal=a,n instanceof Error?(this.data=n.toString(),this.error=n):this.data=n}}function yr(e){return e!=null&&typeof e.status=="number"&&typeof e.statusText=="string"&&typeof e.internal=="boolean"&&"data"in e}const Pt=["post","put","patch","delete"],vr=new Set(Pt),Rr=["get",...Pt],br=new Set(Rr),wr=new Set([301,302,303,307,308]),Dr=new Set([307,308]),Ve={state:"idle",location:void 0,formMethod:void 0,formAction:void 0,formEncType:void 0,formData:void 0,json:void 0,text:void 0},xr={state:"idle",data:void 0,formMethod:void 0,formAction:void 0,formEncType:void 0,formData:void 0,json:void 0,text:void 0},De={state:"unblocked",proceed:void 0,reset:void 0,location:void 0},Ft=/^(?:[a-z][a-z0-9+.-]*:|\/\/)/i,Mr=e=>({hasErrorBoundary:!!e.hasErrorBoundary});function Or(e){const t=e.window?e.window:typeof window<"u"?window:void 0,r=typeof t<"u"&&typeof t.document<"u"&&typeof t.document.createElement<"u",n=!r;U(e.routes.length>0,"You must provide a non-empty routes array to createRouter");let a;if(e.mapRouteProperties)a=e.mapRouteProperties;else if(e.detectErrorBoundary){let o=e.detectErrorBoundary;a=i=>({hasErrorBoundary:o(i)})}else a=Mr;let s={},d=Ye(e.routes,a,void 0,s),u,f=e.basename||"/",h=T({v7_normalizeFormMethod:!1,v7_prependBasename:!1},e.future),m=null,M=new Set,D=null,C=null,E=null,x=e.hydrationData!=null,v=Se(d,e.history.location,f),S=null;if(v==null){let o=V(404,{pathname:e.history.location.pathname}),{matches:i,route:c}=bt(d);v=i,S={[c.id]:o}}let B=!v.some(o=>o.route.lazy)&&(!v.some(o=>o.route.loader)||e.hydrationData!=null),K,l={historyAction:e.history.action,location:e.history.location,matches:v,initialized:B,navigation:Ve,restoreScrollPosition:e.hydrationData!=null?!1:null,preventScrollReset:!1,revalidation:"idle",loaderData:e.hydrationData&&e.hydrationData.loaderData||{},actionData:e.hydrationData&&e.hydrationData.actionData||null,errors:e.hydrationData&&e.hydrationData.errors||S,fetchers:new Map,blockers:new Map},L=H.Pop,$=!1,j,Y=!1,q=!1,X=[],ce=[],O=new Map,Ue=0,ye=-1,ue=new Map,Q=new Set,fe=new Map,ae=new Map,oe=new Map,Oe=!1;function Ct(){return m=e.history.listen(o=>{let{action:i,location:c,delta:p}=o;if(Oe){Oe=!1;return}ge(oe.size===0||p!=null,"You are trying to use a blocker on a POP navigation to a location that was not created by @remix-run/router. This will fail silently in production. This can happen if you are navigating outside the router via `window.history.pushState`/`window.location.hash` instead of using router navigation APIs.  This can also happen if you are using createHashRouter and the user manually changes the URL.");let w=it({currentLocation:l.location,nextLocation:c,historyAction:i});if(w&&p!=null){Oe=!0,e.history.go(p*-1),Ce(w,{state:"blocked",location:c,proceed(){Ce(w,{state:"proceeding",proceed:void 0,reset:void 0,location:c}),e.history.go(p)},reset(){let R=new Map(l.blockers);R.set(w,De),_({blockers:R})}});return}return te(i,c)}),l.initialized||te(H.Pop,l.location),K}function jt(){m&&m(),M.clear(),j&&j.abort(),l.fetchers.forEach((o,i)=>Be(i)),l.blockers.forEach((o,i)=>ot(i))}function It(o){return M.add(o),()=>M.delete(o)}function _(o){l=T({},l,o),M.forEach(i=>i(l))}function ve(o,i){var c,p;let w=l.actionData!=null&&l.navigation.formMethod!=null&&J(l.navigation.formMethod)&&l.navigation.state==="loading"&&((c=o.state)==null?void 0:c._isRedirect)!==!0,R;i.actionData?Object.keys(i.actionData).length>0?R=i.actionData:R=null:w?R=l.actionData:R=null;let b=i.loaderData?Rt(l.loaderData,i.loaderData,i.matches||[],i.errors):l.loaderData,y=l.blockers;y.size>0&&(y=new Map(y),y.forEach((A,W)=>y.set(W,De)));let g=$===!0||l.navigation.formMethod!=null&&J(l.navigation.formMethod)&&((p=o.state)==null?void 0:p._isRedirect)!==!0;u&&(d=u,u=void 0),Y||L===H.Pop||(L===H.Push?e.history.push(o,o.state):L===H.Replace&&e.history.replace(o,o.state)),_(T({},i,{actionData:R,loaderData:b,historyAction:L,location:o,initialized:!0,navigation:Ve,revalidation:"idle",restoreScrollPosition:dt(o,i.matches||l.matches),preventScrollReset:g,blockers:y})),L=H.Pop,$=!1,Y=!1,q=!1,X=[],ce=[]}async function Ze(o,i){if(typeof o=="number"){e.history.go(o);return}let c=Ge(l.location,l.matches,f,h.v7_prependBasename,o,i==null?void 0:i.fromRouteId,i==null?void 0:i.relative),{path:p,submission:w,error:R}=ht(h.v7_normalizeFormMethod,!1,c,i),b=l.location,y=Fe(l.location,p,i&&i.state);y=T({},y,e.history.encodeLocation(y));let g=i&&i.replace!=null?i.replace:void 0,A=H.Push;g===!0?A=H.Replace:g===!1||w!=null&&J(w.formMethod)&&w.formAction===l.location.pathname+l.location.search&&(A=H.Replace);let W=i&&"preventScrollReset"in i?i.preventScrollReset===!0:void 0,P=it({currentLocation:b,nextLocation:y,historyAction:A});if(P){Ce(P,{state:"blocked",location:y,proceed(){Ce(P,{state:"proceeding",proceed:void 0,reset:void 0,location:y}),Ze(o,i)},reset(){let I=new Map(l.blockers);I.set(P,De),_({blockers:I})}});return}return await te(A,y,{submission:w,pendingError:R,preventScrollReset:W,replace:i&&i.replace})}function zt(){if(Ne(),_({revalidation:"loading"}),l.navigation.state!=="submitting"){if(l.navigation.state==="idle"){te(l.historyAction,l.location,{startUninterruptedRevalidation:!0});return}te(L||l.historyAction,l.navigation.location,{overrideNavigation:l.navigation})}}async function te(o,i,c){j&&j.abort(),j=null,L=o,Y=(c&&c.startUninterruptedRevalidation)===!0,$t(l.location,l.matches),$=(c&&c.preventScrollReset)===!0;let p=u||d,w=c&&c.overrideNavigation,R=Se(p,i,f);if(!R){let I=V(404,{pathname:i.pathname}),{matches:k,route:ie}=bt(p);_e(),ve(i,{matches:k,loaderData:{},errors:{[ie.id]:I}});return}if(l.initialized&&!q&&Fr(l.location,i)&&!(c&&c.submission&&J(c.submission.formMethod))){ve(i,{matches:R});return}j=new AbortController;let b=Me(e.history,i,j.signal,c&&c.submission),y,g;if(c&&c.pendingError)g={[me(R).route.id]:c.pendingError};else if(c&&c.submission&&J(c.submission.formMethod)){let I=await Ot(b,i,c.submission,R,{replace:c.replace});if(I.shortCircuited)return;y=I.pendingActionData,g=I.pendingActionError,w=Ie(i,c.submission),b=new Request(b.url,{signal:b.signal})}let{shortCircuited:A,loaderData:W,errors:P}=await Nt(b,i,R,w,c&&c.submission,c&&c.fetcherSubmission,c&&c.replace,y,g);A||(j=null,ve(i,T({matches:R},y?{actionData:y}:{},{loaderData:W,errors:P})))}async function Ot(o,i,c,p,w){w===void 0&&(w={}),Ne();let R=jr(i,c);_({navigation:R});let b,y=Xe(p,i);if(!y.route.action&&!y.route.lazy)b={type:z.error,error:V(405,{method:o.method,pathname:i.pathname,routeId:y.route.id})};else if(b=await xe("action",o,y,p,s,a,f),o.signal.aborted)return{shortCircuited:!0};if(pe(b)){let g;return w&&w.replace!=null?g=w.replace:g=b.location===l.location.pathname+l.location.search,await Re(l,b,{submission:c,replace:g}),{shortCircuited:!0}}if(Pe(b)){let g=me(p,y.route.id);return(w&&w.replace)!==!0&&(L=H.Push),{pendingActionData:{},pendingActionError:{[g.route.id]:b.error}}}if(de(b))throw V(400,{type:"defer-action"});return{pendingActionData:{[y.route.id]:b.data}}}async function Nt(o,i,c,p,w,R,b,y,g){let A=p||Ie(i,w),W=w||R||xt(A),P=u||d,[I,k]=mt(e.history,l,c,W,i,q,X,ce,fe,Q,P,f,y,g);if(_e(F=>!(c&&c.some(G=>G.route.id===F))||I&&I.some(G=>G.route.id===F)),ye=++Ue,I.length===0&&k.length===0){let F=nt();return ve(i,T({matches:c,loaderData:{},errors:g||null},y?{actionData:y}:{},F?{fetchers:new Map(l.fetchers)}:{})),{shortCircuited:!0}}if(!Y){k.forEach(G=>{let ne=l.fetchers.get(G.key),Ke=Ee(void 0,ne?ne.data:void 0);l.fetchers.set(G.key,Ke)});let F=y||l.actionData;_(T({navigation:A},F?Object.keys(F).length===0?{actionData:null}:{actionData:F}:{},k.length>0?{fetchers:new Map(l.fetchers)}:{}))}k.forEach(F=>{O.has(F.key)&&re(F.key),F.controller&&O.set(F.key,F.controller)});let ie=()=>k.forEach(F=>re(F.key));j&&j.signal.addEventListener("abort",ie);let{results:le,loaderResults:be,fetcherResults:ke}=await tt(l.matches,c,I,k,o);if(o.signal.aborted)return{shortCircuited:!0};j&&j.signal.removeEventListener("abort",ie),k.forEach(F=>O.delete(F.key));let Z=wt(le);if(Z){if(Z.idx>=I.length){let F=k[Z.idx-I.length].key;Q.add(F)}return await Re(l,Z.result,{replace:b}),{shortCircuited:!0}}let{loaderData:ee,errors:je}=vt(l,c,I,be,g,k,ke,ae);ae.forEach((F,G)=>{F.subscribe(ne=>{(ne||F.done)&&ae.delete(G)})});let He=nt(),We=at(ye),$e=He||We||k.length>0;return T({loaderData:ee,errors:je},$e?{fetchers:new Map(l.fetchers)}:{})}function et(o){return l.fetchers.get(o)||xr}function Bt(o,i,c,p){if(n)throw new Error("router.fetch() was called during the server render, but it shouldn't be. You are likely calling a useFetcher() method in the body of your component. Try moving it to a useEffect or a callback.");O.has(o)&&re(o);let w=u||d,R=Ge(l.location,l.matches,f,h.v7_prependBasename,c,i,p==null?void 0:p.relative),b=Se(w,R,f);if(!b){Te(o,i,V(404,{pathname:R}));return}let{path:y,submission:g,error:A}=ht(h.v7_normalizeFormMethod,!0,R,p);if(A){Te(o,i,A);return}let W=Xe(b,y);if($=(p&&p.preventScrollReset)===!0,g&&J(g.formMethod)){_t(o,i,y,W,b,g);return}fe.set(o,{routeId:i,path:y}),kt(o,i,y,W,b,g)}async function _t(o,i,c,p,w,R){if(Ne(),fe.delete(o),!p.route.action&&!p.route.lazy){let N=V(405,{method:R.formMethod,pathname:c,routeId:i});Te(o,i,N);return}let b=l.fetchers.get(o),y=Ir(R,b);l.fetchers.set(o,y),_({fetchers:new Map(l.fetchers)});let g=new AbortController,A=Me(e.history,c,g.signal,R);O.set(o,g);let W=Ue,P=await xe("action",A,p,w,s,a,f);if(A.signal.aborted){O.get(o)===g&&O.delete(o);return}if(pe(P))if(O.delete(o),ye>W){let N=he(void 0);l.fetchers.set(o,N),_({fetchers:new Map(l.fetchers)});return}else{Q.add(o);let N=Ee(R);return l.fetchers.set(o,N),_({fetchers:new Map(l.fetchers)}),Re(l,P,{submission:R,isFetchActionRedirect:!0})}if(Pe(P)){Te(o,i,P.error);return}if(de(P))throw V(400,{type:"defer-action"});let I=l.navigation.location||l.location,k=Me(e.history,I,g.signal),ie=u||d,le=l.navigation.state!=="idle"?Se(ie,l.navigation.location,f):l.matches;U(le,"Didn't find any matches after fetcher action");let be=++Ue;ue.set(o,be);let ke=Ee(R,P.data);l.fetchers.set(o,ke);let[Z,ee]=mt(e.history,l,le,R,I,q,X,ce,fe,Q,ie,f,{[p.route.id]:P.data},void 0);ee.filter(N=>N.key!==o).forEach(N=>{let we=N.key,st=l.fetchers.get(we),qt=Ee(void 0,st?st.data:void 0);l.fetchers.set(we,qt),O.has(we)&&re(we),N.controller&&O.set(we,N.controller)}),_({fetchers:new Map(l.fetchers)});let je=()=>ee.forEach(N=>re(N.key));g.signal.addEventListener("abort",je);let{results:He,loaderResults:We,fetcherResults:$e}=await tt(l.matches,le,Z,ee,k);if(g.signal.aborted)return;g.signal.removeEventListener("abort",je),ue.delete(o),O.delete(o),ee.forEach(N=>O.delete(N.key));let F=wt(He);if(F){if(F.idx>=Z.length){let N=ee[F.idx-Z.length].key;Q.add(N)}return Re(l,F.result)}let{loaderData:G,errors:ne}=vt(l,l.matches,Z,We,void 0,ee,$e,ae);if(l.fetchers.has(o)){let N=he(P.data);l.fetchers.set(o,N)}let Ke=at(be);l.navigation.state==="loading"&&be>ye?(U(L,"Expected pending action"),j&&j.abort(),ve(l.navigation.location,{matches:le,loaderData:G,errors:ne,fetchers:new Map(l.fetchers)})):(_(T({errors:ne,loaderData:Rt(l.loaderData,G,le,ne)},Ke||ee.length>0?{fetchers:new Map(l.fetchers)}:{})),q=!1)}async function kt(o,i,c,p,w,R){let b=l.fetchers.get(o),y=Ee(R,b?b.data:void 0);l.fetchers.set(o,y),_({fetchers:new Map(l.fetchers)});let g=new AbortController,A=Me(e.history,c,g.signal);O.set(o,g);let W=Ue,P=await xe("loader",A,p,w,s,a,f);if(de(P)&&(P=await Tt(P,A.signal,!0)||P),O.get(o)===g&&O.delete(o),A.signal.aborted)return;if(pe(P))if(ye>W){let k=he(void 0);l.fetchers.set(o,k),_({fetchers:new Map(l.fetchers)});return}else{Q.add(o),await Re(l,P);return}if(Pe(P)){let k=me(l.matches,i);l.fetchers.delete(o),_({fetchers:new Map(l.fetchers),errors:{[k.route.id]:P.error}});return}U(!de(P),"Unhandled fetcher deferred data");let I=he(P.data);l.fetchers.set(o,I),_({fetchers:new Map(l.fetchers)})}async function Re(o,i,c){let{submission:p,replace:w,isFetchActionRedirect:R}=c===void 0?{}:c;i.revalidate&&(q=!0);let b=Fe(o.location,i.location,T({_isRedirect:!0},R?{_isFetchActionRedirect:!0}:{}));if(U(b,"Expected a location on the redirect navigation"),r){let A=!1;if(i.reloadDocument)A=!0;else if(Ft.test(i.location)){const W=e.history.createURL(i.location);A=W.origin!==t.location.origin||ze(W.pathname,f)==null}if(A){w?t.location.replace(i.location):t.location.assign(i.location);return}}j=null;let y=w===!0?H.Replace:H.Push,g=p||xt(o.navigation);if(Dr.has(i.status)&&g&&J(g.formMethod))await te(y,b,{submission:T({},g,{formAction:i.location}),preventScrollReset:$});else if(R)await te(y,b,{overrideNavigation:Ie(b),fetcherSubmission:g,preventScrollReset:$});else{let A=Ie(b,g);await te(y,b,{overrideNavigation:A,preventScrollReset:$})}}async function tt(o,i,c,p,w){let R=await Promise.all([...c.map(g=>xe("loader",w,g,i,s,a,f)),...p.map(g=>g.matches&&g.match&&g.controller?xe("loader",Me(e.history,g.path,g.controller.signal),g.match,g.matches,s,a,f):{type:z.error,error:V(404,{pathname:g.path})})]),b=R.slice(0,c.length),y=R.slice(c.length);return await Promise.all([Dt(o,c,b,b.map(()=>w.signal),!1,l.loaderData),Dt(o,p.map(g=>g.match),y,p.map(g=>g.controller?g.controller.signal:null),!0)]),{results:R,loaderResults:b,fetcherResults:y}}function Ne(){q=!0,X.push(..._e()),fe.forEach((o,i)=>{O.has(i)&&(ce.push(i),re(i))})}function Te(o,i,c){let p=me(l.matches,i);Be(o),_({errors:{[p.route.id]:c},fetchers:new Map(l.fetchers)})}function Be(o){let i=l.fetchers.get(o);O.has(o)&&!(i&&i.state==="loading"&&ue.has(o))&&re(o),fe.delete(o),ue.delete(o),Q.delete(o),l.fetchers.delete(o)}function re(o){let i=O.get(o);U(i,"Expected fetch controller: "+o),i.abort(),O.delete(o)}function rt(o){for(let i of o){let c=et(i),p=he(c.data);l.fetchers.set(i,p)}}function nt(){let o=[],i=!1;for(let c of Q){let p=l.fetchers.get(c);U(p,"Expected fetcher: "+c),p.state==="loading"&&(Q.delete(c),o.push(c),i=!0)}return rt(o),i}function at(o){let i=[];for(let[c,p]of ue)if(p<o){let w=l.fetchers.get(c);U(w,"Expected fetcher: "+c),w.state==="loading"&&(re(c),ue.delete(c),i.push(c))}return rt(i),i.length>0}function Ht(o,i){let c=l.blockers.get(o)||De;return oe.get(o)!==i&&oe.set(o,i),c}function ot(o){l.blockers.delete(o),oe.delete(o)}function Ce(o,i){let c=l.blockers.get(o)||De;U(c.state==="unblocked"&&i.state==="blocked"||c.state==="blocked"&&i.state==="blocked"||c.state==="blocked"&&i.state==="proceeding"||c.state==="blocked"&&i.state==="unblocked"||c.state==="proceeding"&&i.state==="unblocked","Invalid blocker state transition: "+c.state+" -> "+i.state);let p=new Map(l.blockers);p.set(o,i),_({blockers:p})}function it(o){let{currentLocation:i,nextLocation:c,historyAction:p}=o;if(oe.size===0)return;oe.size>1&&ge(!1,"A router only supports one blocker at a time");let w=Array.from(oe.entries()),[R,b]=w[w.length-1],y=l.blockers.get(R);if(!(y&&y.state==="proceeding")&&b({currentLocation:i,nextLocation:c,historyAction:p}))return R}function _e(o){let i=[];return ae.forEach((c,p)=>{(!o||o(p))&&(c.cancel(),i.push(p),ae.delete(p))}),i}function Wt(o,i,c){if(D=o,E=i,C=c||null,!x&&l.navigation===Ve){x=!0;let p=dt(l.location,l.matches);p!=null&&_({restoreScrollPosition:p})}return()=>{D=null,E=null,C=null}}function lt(o,i){return C&&C(o,i.map(p=>Cr(p,l.loaderData)))||o.key}function $t(o,i){if(D&&E){let c=lt(o,i);D[c]=E()}}function dt(o,i){if(D){let c=lt(o,i),p=D[c];if(typeof p=="number")return p}return null}function Kt(o){s={},u=Ye(o,a,void 0,s)}return K={get basename(){return f},get state(){return l},get routes(){return d},initialize:Ct,subscribe:It,enableScrollRestoration:Wt,navigate:Ze,fetch:Bt,revalidate:zt,createHref:o=>e.history.createHref(o),encodeLocation:o=>e.history.encodeLocation(o),getFetcher:et,deleteFetcher:Be,dispose:jt,getBlocker:Ht,deleteBlocker:ot,_internalFetchControllers:O,_internalActiveDeferreds:ae,_internalSetRoutes:Kt},K}function Er(e){return e!=null&&("formData"in e&&e.formData!=null||"body"in e&&e.body!==void 0)}function Ge(e,t,r,n,a,s,d){let u,f;if(s!=null&&d!=="path"){u=[];for(let m of t)if(u.push(m),m.route.id===s){f=m;break}}else u=t,f=t[t.length-1];let h=hr(a||".",St(u).map(m=>m.pathnameBase),ze(e.pathname,r)||e.pathname,d==="path");return a==null&&(h.search=e.search,h.hash=e.hash),(a==null||a===""||a===".")&&f&&f.route.index&&!Qe(h.search)&&(h.search=h.search?h.search.replace(/^\?/,"?index&"):"?index"),n&&r!=="/"&&(h.pathname=h.pathname==="/"?r:Le([r,h.pathname])),Ae(h)}function ht(e,t,r,n){if(!n||!Er(n))return{path:r};if(n.formMethod&&!Tr(n.formMethod))return{path:r,error:V(405,{method:n.formMethod})};let a=()=>({path:r,error:V(400,{type:"invalid-body"})}),s=n.formMethod||"get",d=e?s.toUpperCase():s.toLowerCase(),u=Ut(r);if(n.body!==void 0){if(n.formEncType==="text/plain"){if(!J(d))return a();let D=typeof n.body=="string"?n.body:n.body instanceof FormData||n.body instanceof URLSearchParams?Array.from(n.body.entries()).reduce((C,E)=>{let[x,v]=E;return""+C+x+"="+v+`
`},""):String(n.body);return{path:r,submission:{formMethod:d,formAction:u,formEncType:n.formEncType,formData:void 0,json:void 0,text:D}}}else if(n.formEncType==="application/json"){if(!J(d))return a();try{let D=typeof n.body=="string"?JSON.parse(n.body):n.body;return{path:r,submission:{formMethod:d,formAction:u,formEncType:n.formEncType,formData:void 0,json:D,text:void 0}}}catch{return a()}}}U(typeof FormData=="function","FormData is not available in this environment");let f,h;if(n.formData)f=Je(n.formData),h=n.formData;else if(n.body instanceof FormData)f=Je(n.body),h=n.body;else if(n.body instanceof URLSearchParams)f=n.body,h=yt(f);else if(n.body==null)f=new URLSearchParams,h=new FormData;else try{f=new URLSearchParams(n.body),h=yt(f)}catch{return a()}let m={formMethod:d,formAction:u,formEncType:n&&n.formEncType||"application/x-www-form-urlencoded",formData:h,json:void 0,text:void 0};if(J(m.formMethod))return{path:r,submission:m};let M=se(r);return t&&M.search&&Qe(M.search)&&f.append("index",""),M.search="?"+f,{path:Ae(M),submission:m}}function Sr(e,t){let r=e;if(t){let n=e.findIndex(a=>a.route.id===t);n>=0&&(r=e.slice(0,n))}return r}function mt(e,t,r,n,a,s,d,u,f,h,m,M,D,C){let E=C?Object.values(C)[0]:D?Object.values(D)[0]:void 0,x=e.createURL(t.location),v=e.createURL(a),S=C?Object.keys(C)[0]:void 0,K=Sr(r,S).filter((L,$)=>{if(L.route.lazy)return!0;if(L.route.loader==null)return!1;if(Lr(t.loaderData,t.matches[$],L)||d.some(q=>q===L.route.id))return!0;let j=t.matches[$],Y=L;return pt(L,T({currentUrl:x,currentParams:j.params,nextUrl:v,nextParams:Y.params},n,{actionResult:E,defaultShouldRevalidate:s||x.pathname+x.search===v.pathname+v.search||x.search!==v.search||At(j,Y)}))}),l=[];return f.forEach((L,$)=>{if(!r.some(ce=>ce.route.id===L.routeId))return;let j=Se(m,L.path,M);if(!j){l.push({key:$,routeId:L.routeId,path:L.path,matches:null,match:null,controller:null});return}let Y=t.fetchers.get($),q=Xe(j,L.path),X=!1;h.has($)?X=!1:u.includes($)?X=!0:Y&&Y.state!=="idle"&&Y.data===void 0?X=s:X=pt(q,T({currentUrl:x,currentParams:t.matches[t.matches.length-1].params,nextUrl:v,nextParams:r[r.length-1].params},n,{actionResult:E,defaultShouldRevalidate:s})),X&&l.push({key:$,routeId:L.routeId,path:L.path,matches:j,match:q,controller:new AbortController})}),[K,l]}function Lr(e,t,r){let n=!t||r.route.id!==t.route.id,a=e[r.route.id]===void 0;return n||a}function At(e,t){let r=e.route.path;return e.pathname!==t.pathname||r!=null&&r.endsWith("*")&&e.params["*"]!==t.params["*"]}function pt(e,t){if(e.route.shouldRevalidate){let r=e.route.shouldRevalidate(t);if(typeof r=="boolean")return r}return t.defaultShouldRevalidate}async function gt(e,t,r){if(!e.lazy)return;let n=await e.lazy();if(!e.lazy)return;let a=r[e.id];U(a,"No route found in manifest");let s={};for(let d in n){let f=a[d]!==void 0&&d!=="hasErrorBoundary";ge(!f,'Route "'+a.id+'" has a static property "'+d+'" defined but its lazy function is also returning a value for this property. '+('The lazy route property "'+d+'" will be ignored.')),!f&&!Gt.has(d)&&(s[d]=n[d])}Object.assign(a,s),Object.assign(a,T({},t(a),{lazy:void 0}))}async function xe(e,t,r,n,a,s,d,u){u===void 0&&(u={});let f,h,m,M=E=>{let x,v=new Promise((S,B)=>x=B);return m=()=>x(),t.signal.addEventListener("abort",m),Promise.race([E({request:t,params:r.params,context:u.requestContext}),v])};try{let E=r.route[e];if(r.route.lazy)if(E)h=(await Promise.all([M(E),gt(r.route,s,a)]))[0];else if(await gt(r.route,s,a),E=r.route[e],E)h=await M(E);else if(e==="action"){let x=new URL(t.url),v=x.pathname+x.search;throw V(405,{method:t.method,pathname:v,routeId:r.route.id})}else return{type:z.data,data:void 0};else if(E)h=await M(E);else{let x=new URL(t.url),v=x.pathname+x.search;throw V(404,{pathname:v})}U(h!==void 0,"You defined "+(e==="action"?"an action":"a loader")+" for route "+('"'+r.route.id+"\" but didn't return anything from your `"+e+"` ")+"function. Please return a value or `null`.")}catch(E){f=z.error,h=E}finally{m&&t.signal.removeEventListener("abort",m)}if(Ur(h)){let E=h.status;if(wr.has(E)){let S=h.headers.get("Location");if(U(S,"Redirects returned/thrown from loaders/actions must have a Location header"),!Ft.test(S))S=Ge(new URL(t.url),n.slice(0,n.indexOf(r)+1),d,!0,S);else if(!u.isStaticRequest){let B=new URL(t.url),K=S.startsWith("//")?new URL(B.protocol+S):new URL(S),l=ze(K.pathname,d)!=null;K.origin===B.origin&&l&&(S=K.pathname+K.search+K.hash)}if(u.isStaticRequest)throw h.headers.set("Location",S),h;return{type:z.redirect,status:E,location:S,revalidate:h.headers.get("X-Remix-Revalidate")!==null,reloadDocument:h.headers.get("X-Remix-Reload-Document")!==null}}if(u.isRouteRequest)throw{type:f===z.error?z.error:z.data,response:h};let x,v=h.headers.get("Content-Type");return v&&/\bapplication\/json\b/.test(v)?x=await h.json():x=await h.text(),f===z.error?{type:f,error:new Lt(E,h.statusText,x),headers:h.headers}:{type:z.data,data:x,statusCode:h.status,headers:h.headers}}if(f===z.error)return{type:f,error:h};if(Ar(h)){var D,C;return{type:z.deferred,deferredData:h,statusCode:(D=h.init)==null?void 0:D.status,headers:((C=h.init)==null?void 0:C.headers)&&new Headers(h.init.headers)}}return{type:z.data,data:h}}function Me(e,t,r,n){let a=e.createURL(Ut(t)).toString(),s={signal:r};if(n&&J(n.formMethod)){let{formMethod:d,formEncType:u}=n;s.method=d.toUpperCase(),u==="application/json"?(s.headers=new Headers({"Content-Type":u}),s.body=JSON.stringify(n.json)):u==="text/plain"?s.body=n.text:u==="application/x-www-form-urlencoded"&&n.formData?s.body=Je(n.formData):s.body=n.formData}return new Request(a,s)}function Je(e){let t=new URLSearchParams;for(let[r,n]of e.entries())t.append(r,typeof n=="string"?n:n.name);return t}function yt(e){let t=new FormData;for(let[r,n]of e.entries())t.append(r,n);return t}function Pr(e,t,r,n,a){let s={},d=null,u,f=!1,h={};return r.forEach((m,M)=>{let D=t[M].route.id;if(U(!pe(m),"Cannot handle redirect results in processLoaderData"),Pe(m)){let C=me(e,D),E=m.error;n&&(E=Object.values(n)[0],n=void 0),d=d||{},d[C.route.id]==null&&(d[C.route.id]=E),s[D]=void 0,f||(f=!0,u=yr(m.error)?m.error.status:500),m.headers&&(h[D]=m.headers)}else de(m)?(a.set(D,m.deferredData),s[D]=m.deferredData.data):s[D]=m.data,m.statusCode!=null&&m.statusCode!==200&&!f&&(u=m.statusCode),m.headers&&(h[D]=m.headers)}),n&&(d=n,s[Object.keys(n)[0]]=void 0),{loaderData:s,errors:d,statusCode:u||200,loaderHeaders:h}}function vt(e,t,r,n,a,s,d,u){let{loaderData:f,errors:h}=Pr(t,r,n,a,u);for(let m=0;m<s.length;m++){let{key:M,match:D,controller:C}=s[m];U(d!==void 0&&d[m]!==void 0,"Did not find corresponding fetcher result");let E=d[m];if(!(C&&C.signal.aborted))if(Pe(E)){let x=me(e.matches,D==null?void 0:D.route.id);h&&h[x.route.id]||(h=T({},h,{[x.route.id]:E.error})),e.fetchers.delete(M)}else if(pe(E))U(!1,"Unhandled fetcher revalidation redirect");else if(de(E))U(!1,"Unhandled fetcher deferred data");else{let x=he(E.data);e.fetchers.set(M,x)}}return{loaderData:f,errors:h}}function Rt(e,t,r,n){let a=T({},t);for(let s of r){let d=s.route.id;if(t.hasOwnProperty(d)?t[d]!==void 0&&(a[d]=t[d]):e[d]!==void 0&&s.route.loader&&(a[d]=e[d]),n&&n.hasOwnProperty(d))break}return a}function me(e,t){return(t?e.slice(0,e.findIndex(n=>n.route.id===t)+1):[...e]).reverse().find(n=>n.route.hasErrorBoundary===!0)||e[0]}function bt(e){let t=e.find(r=>r.index||!r.path||r.path==="/")||{id:"__shim-error-route__"};return{matches:[{params:{},pathname:"",pathnameBase:"",route:t}],route:t}}function V(e,t){let{pathname:r,routeId:n,method:a,type:s}=t===void 0?{}:t,d="Unknown Server Error",u="Unknown @remix-run/router error";return e===400?(d="Bad Request",a&&r&&n?u="You made a "+a+' request to "'+r+'" but '+('did not provide a `loader` for route "'+n+'", ')+"so there is no way to handle the request.":s==="defer-action"?u="defer() is not supported in actions":s==="invalid-body"&&(u="Unable to encode submission body")):e===403?(d="Forbidden",u='Route "'+n+'" does not match URL "'+r+'"'):e===404?(d="Not Found",u='No route matches URL "'+r+'"'):e===405&&(d="Method Not Allowed",a&&r&&n?u="You made a "+a.toUpperCase()+' request to "'+r+'" but '+('did not provide an `action` for route "'+n+'", ')+"so there is no way to handle the request.":a&&(u='Invalid request method "'+a.toUpperCase()+'"')),new Lt(e||500,d,new Error(u),!0)}function wt(e){for(let t=e.length-1;t>=0;t--){let r=e[t];if(pe(r))return{result:r,idx:t}}}function Ut(e){let t=typeof e=="string"?se(e):e;return Ae(T({},t,{hash:""}))}function Fr(e,t){return e.pathname!==t.pathname||e.search!==t.search?!1:e.hash===""?t.hash!=="":e.hash===t.hash?!0:t.hash!==""}function de(e){return e.type===z.deferred}function Pe(e){return e.type===z.error}function pe(e){return(e&&e.type)===z.redirect}function Ar(e){let t=e;return t&&typeof t=="object"&&typeof t.data=="object"&&typeof t.subscribe=="function"&&typeof t.cancel=="function"&&typeof t.resolveData=="function"}function Ur(e){return e!=null&&typeof e.status=="number"&&typeof e.statusText=="string"&&typeof e.headers=="object"&&typeof e.body<"u"}function Tr(e){return br.has(e.toLowerCase())}function J(e){return vr.has(e.toLowerCase())}async function Dt(e,t,r,n,a,s){for(let d=0;d<r.length;d++){let u=r[d],f=t[d];if(!f)continue;let h=e.find(M=>M.route.id===f.route.id),m=h!=null&&!At(h,f)&&(s&&s[f.route.id])!==void 0;if(de(u)&&(a||m)){let M=n[d];U(M,"Expected an AbortSignal for revalidating fetcher deferred result"),await Tt(u,M,a).then(D=>{D&&(r[d]=D||r[d])})}}}async function Tt(e,t,r){if(r===void 0&&(r=!1),!await e.deferredData.resolveData(t)){if(r)try{return{type:z.data,data:e.deferredData.unwrappedData}}catch(a){return{type:z.error,error:a}}return{type:z.data,data:e.deferredData.data}}}function Qe(e){return new URLSearchParams(e).getAll("index").some(t=>t==="")}function Cr(e,t){let{route:r,pathname:n,params:a}=e;return{id:r.id,pathname:n,params:a,data:t[r.id],handle:r.handle}}function Xe(e,t){let r=typeof t=="string"?se(t).search:t.search;if(e[e.length-1].route.index&&Qe(r||""))return e[e.length-1];let n=St(e);return n[n.length-1]}function xt(e){let{formMethod:t,formAction:r,formEncType:n,text:a,formData:s,json:d}=e;if(!(!t||!r||!n)){if(a!=null)return{formMethod:t,formAction:r,formEncType:n,formData:void 0,json:void 0,text:a};if(s!=null)return{formMethod:t,formAction:r,formEncType:n,formData:s,json:void 0,text:void 0};if(d!==void 0)return{formMethod:t,formAction:r,formEncType:n,formData:void 0,json:d,text:void 0}}}function Ie(e,t){return t?{state:"loading",location:e,formMethod:t.formMethod,formAction:t.formAction,formEncType:t.formEncType,formData:t.formData,json:t.json,text:t.text}:{state:"loading",location:e,formMethod:void 0,formAction:void 0,formEncType:void 0,formData:void 0,json:void 0,text:void 0}}function jr(e,t){return{state:"submitting",location:e,formMethod:t.formMethod,formAction:t.formAction,formEncType:t.formEncType,formData:t.formData,json:t.json,text:t.text}}function Ee(e,t){return e?{state:"loading",formMethod:e.formMethod,formAction:e.formAction,formEncType:e.formEncType,formData:e.formData,json:e.json,text:e.text,data:t," _hasFetcherDoneAnything ":!0}:{state:"loading",formMethod:void 0,formAction:void 0,formEncType:void 0,formData:void 0,json:void 0,text:void 0,data:t," _hasFetcherDoneAnything ":!0}}function Ir(e,t){return{state:"submitting",formMethod:e.formMethod,formAction:e.formAction,formEncType:e.formEncType,formData:e.formData,json:e.json,text:e.text,data:t?t.data:void 0," _hasFetcherDoneAnything ":!0}}function he(e){return{state:"idle",formMethod:void 0,formAction:void 0,formEncType:void 0,formData:void 0,json:void 0,text:void 0,data:e," _hasFetcherDoneAnything ":!0}}export{H as A,Lt as E,yr as a,zr as b,Or as c,St as g,U as i,Le as j,Se as m,se as p,hr as r,ze as s};
