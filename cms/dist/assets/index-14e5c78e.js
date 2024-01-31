import{j as e,r,R as we}from"./react-6a422704.js";import{c as Ce}from"./react-dom-81c1590d.js";import{Q as Se,u as G,a as Z,b as Me,d as Re}from"./react-query-5a384075.js";import{c as ke}from"./react-router-dom-8a25b6f8.js";import{N as M,L as R,T as U,a as ee,M as C,B as I,b as $e,C as te,R as V}from"./flowbite-react-95944268.js";import"./react-indiana-drag-scroll-d13610c9.js";import"./firebase-5ecc8669.js";import{i as Ie,g as Te,s as Ae}from"./@firebase-7f7857ef.js";import{u as Ee,a as Fe,R as Pe}from"./react-router-33fe3b2e.js";import{a as A}from"./axios-4a70c6fc.js";import{u as Oe,f as J,g as De,a as ze,b as Ue,c as Le}from"./@tanstack-24ebb7b7.js";import{$ as Ve,a as se,b as ae,c as Be,d as re,e as ne,f as oe,g as le,h as ie,i as de,j as ce,k as He,l as qe,m as me,n as Ke}from"./@radix-ui-e33ec21f.js";import{c as Qe}from"./class-variance-authority-52f2569e.js";import{c as Ye}from"./clsx-1229b3e0.js";import{t as Ge}from"./tailwind-merge-0405149e.js";import{C as _e,a as ue,b as We,A as B,M as Je,c as Xe}from"./lucide-react-9dcfc33d.js";import"./scheduler-765c72db.js";import"./@remix-run-0093eeff.js";import"./react-icons-bbfbd1e9.js";import"./@floating-ui-98aebd19.js";import"./aria-hidden-318acb9e.js";import"./tabbable-6ab97b54.js";import"./idb-81bdbf9b.js";import"./tslib-89fde298.js";import"./@babel-98964cd2.js";import"./react-remove-scroll-b71c63fb.js";import"./react-remove-scroll-bar-0babf0e5.js";import"./react-style-singleton-c9b02a91.js";import"./get-nonce-1ea010b7.js";import"./use-sidecar-1e201c94.js";import"./use-callback-ref-601dd590.js";(function(){const s=document.createElement("link").relList;if(s&&s.supports&&s.supports("modulepreload"))return;for(const o of document.querySelectorAll('link[rel="modulepreload"]'))n(o);new MutationObserver(o=>{for(const l of o)if(l.type==="childList")for(const d of l.addedNodes)d.tagName==="LINK"&&d.rel==="modulepreload"&&n(d)}).observe(document,{childList:!0,subtree:!0});function a(o){const l={};return o.integrity&&(l.integrity=o.integrity),o.referrerPolicy&&(l.referrerPolicy=o.referrerPolicy),o.crossOrigin==="use-credentials"?l.credentials="include":o.crossOrigin==="anonymous"?l.credentials="omit":l.credentials="same-origin",l}function n(o){if(o.ep)return;o.ep=!0;const l=a(o);fetch(o.href,l)}})();const Ze=new Se,et="67a78d53-d47c-4981-b3e4-1843b8591904",H="/notifications",tt="/users",st="/dashboard",q="/messages",K="/payments",at={apiKey:"AIzaSyC59xS5P27qzafvucTLTPkAm6gwrWRfMQ8",authDomain:"momentum-ministry.firebaseapp.com",projectId:"momentum-ministry",storageBucket:"momentum-ministry.appspot.com",messagingSenderId:"649626638654",appId:"1:649626638654:web:2c65617d5ee8ed16cab46f",measurementId:"G-F5RPPG148R"},rt=Ie(at),b=Te(rt);function nt({onSigninClick:t}){const[s,a]=r.useState(!!b.currentUser);return r.useEffect(()=>{b.onAuthStateChanged(n=>{a(!!n)}),b.currentUser?a(!0):a(!1)},[b.currentUser]),s?e.jsx(M.Link,{className:"cursor-pointer",onClick:()=>{b.signOut().then(()=>{alert("Signed out")}).catch(n=>{alert(n)})},children:"Sign Out"}):e.jsx(M.Link,{onClick:t,children:"Sign In"})}function P({onSigninClick:t}){let s=Ee();return e.jsxs(M,{fluid:!0,rounded:!0,className:"w-full",children:[e.jsxs(M.Brand,{href:"/dashboard",children:[e.jsx("img",{src:"assets/logo.png",className:"mr-3 h-6 sm:h-9",alt:"Momentum Logo"}),e.jsx("span",{className:"self-center whitespace-nowrap text-xl font-bold border-b border-gray-100 text-gray-700 hover:bg-gray-50 dark:border-gray-700 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white md:border-0 md:hover:bg-transparent md:hover:text-cyan-700 md:dark:hover:bg-transparent md:dark:hover:text-white",children:"Momentum Dashboard"})]}),e.jsx(M.Toggle,{}),e.jsxs(M.Collapse,{children:[e.jsx(M.Link,{href:q,active:q===s.pathname,children:"Messages"}),e.jsx(M.Link,{href:H,active:H===s.pathname,children:"Notifications"}),e.jsx(M.Link,{href:K,active:K===s.pathname,children:"Transactions"}),e.jsx(nt,{onSigninClick:t})]})]})}const ot=()=>{const t=Fe();return r.useEffect(()=>{t("/messages")},[]),e.jsx("div",{className:"w-full",children:e.jsx(P,{onSigninClick:()=>{}})})};function N({reference:t,value:s,title:a,placeholder:n,id:o,onChange:l}){return e.jsxs("div",{children:[e.jsx("div",{className:"mb-2 block",children:e.jsx(R,{htmlFor:o,value:a})}),e.jsx(U,{id:o,ref:t,placeholder:n,value:s,onChange:l,required:!0})]})}function L({reference:t,value:s,title:a,placeholder:n,id:o,onChange:l}){return e.jsxs("div",{children:[e.jsx("div",{className:"mb-2 block",children:e.jsx(R,{htmlFor:o,value:a})}),e.jsx(ee,{id:o,ref:t,placeholder:n,value:s,onChange:l,required:!0,rows:6})]})}function lt({openModal:t,setOpenModal:s}){const[a,n]=r.useState([]),o=r.useRef(null),l=r.useRef(null),d=r.useRef(null),p=r.useRef(null),c=r.useRef(null),f=G(u=>A.post("/api/v1/messages",u),{onSuccess:()=>{alert("Message posted successfully")}});return e.jsxs(C,{show:t,size:"md",popup:!0,onClose:()=>s(!1),initialFocus:o,children:[e.jsx(C.Header,{}),e.jsx(C.Body,{children:e.jsxs("div",{className:"space-y-6",children:[e.jsx("h3",{className:"text-xl font-medium text-gray-900 dark:text-white",children:"Add Message"}),e.jsx(N,{reference:o,title:"Title",placeholder:"Enter the title of the message",id:"title"}),e.jsx(N,{reference:l,title:"Series",placeholder:"Enter the series of the message",id:"series"}),e.jsx(N,{reference:d,title:"Preacher",placeholder:"Enter the name of the preacher",id:"preacher"}),e.jsx(N,{reference:c,title:"Date",placeholder:"Enter the display date of the message",id:"date"}),e.jsx(N,{reference:p,title:"Image",placeholder:"Enter the image cover url of the message",id:"image"}),a.map((u,v)=>{switch(u.type){case 0:return e.jsx(L,{reference:u.area,title:"Add header",placeholder:"Enter the header of the passage",id:`header${v+1}`});case 1:return e.jsxs("div",{children:[e.jsx(N,{reference:u.text,title:"Add verse",placeholder:"Enter the verse of the passage",id:`verse${v+1}`}),e.jsx(L,{reference:u.area,title:"Add message",placeholder:"Enter the message of the passage",id:`message${v+1}`})]})}}),e.jsxs("div",{className:"w-full flex justify-between",children:[e.jsx(I,{className:"bg-momentum-orange hover:bg-momentum-orange enabled:hover:bg-momentum-orange",onClick:()=>{var u,v,m,x,E;f.mutate({title:((u=o.current)==null?void 0:u.value)??"",series:((v=l.current)==null?void 0:v.value)??"",preacher:((m=d.current)==null?void 0:m.value)??"",thumbnail:((x=p.current)==null?void 0:x.value)??"",date:((E=c.current)==null?void 0:E.value)??"",passages:a.map((y,k)=>{var j,S,O;switch(y.type){case 0:return{header:((j=y.area.current)==null?void 0:j.value)??"",type:"header",order:k};case 1:return{verse:((S=y.text.current)==null?void 0:S.value)??"",message:((O=y.area.current)==null?void 0:O.value)??"",type:"message",order:k}}})}),s(!1)},children:"Submit"}),e.jsxs("div",{className:"flex",children:[e.jsx(I,{className:"bg-momentum-orange hover:bg-momentum-orange active:bg-momentum-orange enabled:hover:bg-momentum-orange",onClick:()=>{n([...a,{type:0,text:r.createRef(),area:r.createRef()}]),console.log(a)},children:"Header"}),e.jsxs("div",{className:"flex",children:[e.jsx(I,{className:"ml-1 bg-momentum-orange hover:bg-momentum-orange enabled:hover:bg-momentum-orange",onClick:()=>{n([...a,{type:1,text:r.createRef(),area:r.createRef()}]),console.log(a)},children:"Message"}),a.length>0?e.jsx(I,{className:"ml-1 rounded-full bg-momentum-orange hover:bg-momentum-orange enabled:hover:bg-momentum-orange",onClick:()=>{n([...a.slice(0,a.length-1)]),console.log(a)},children:e.jsx("svg",{className:"w-4 h-4 rounded-full text-white-800 dark:text-white","aria-hidden":"true",xmlns:"http://www.w3.org/2000/svg",fill:"none",viewBox:"0 0 24 24",children:e.jsx("path",{stroke:"currentColor","stroke-linecap":"round","stroke-linejoin":"round","stroke-width":"2",d:"M3 9h13a5 5 0 0 1 0 10H7M3 9l4-4M3 9l4 4"})})}):e.jsx(e.Fragment,{})]})]})]})]})})]})}function _({openModal:t,setOpenModal:s}){const a=r.useRef(null),n=r.useRef(null);return e.jsxs(C,{show:t,size:"md",popup:!0,onClose:()=>s(!1),initialFocus:a,children:[e.jsx(C.Header,{}),e.jsx(C.Body,{children:e.jsxs("div",{className:"space-y-6",children:[e.jsx("h3",{className:"text-xl font-medium text-gray-900 dark:text-white",children:"Sign in to your account"}),e.jsxs("div",{children:[e.jsx("div",{className:"mb-2 block",children:e.jsx(R,{htmlFor:"email",value:"Your email"})}),e.jsx(U,{id:"email",ref:a,placeholder:"name@company.com",required:!0})]}),e.jsxs("div",{children:[e.jsx("div",{className:"mb-2 block",children:e.jsx(R,{htmlFor:"password",value:"Your password"})}),e.jsx(U,{id:"password",ref:n,type:"password",required:!0})]}),e.jsx("div",{className:"w-full",children:e.jsx(I,{className:"bg-momentum-orange hover:bg-momentum-orange enabled:hover:bg-momentum-orange",onClick:()=>{var o,l;Ae(b,(o=a.current)==null?void 0:o.value,(l=n.current)==null?void 0:l.value).then(d=>{const p=d.user;s(!1),alert(`Signed in as ${p==null?void 0:p.email}`)}).catch(d=>{const p=d.code,c=d.message;alert(`Error: ${p} ${c}`)})},children:"Log in to your account"})})]})})]})}function it({message:t,openModal:s,setOpenModal:a}){const[n,o]=r.useState(""),[l,d]=r.useState(""),[p,c]=r.useState(""),[f,u]=r.useState(""),[v,m]=r.useState(""),[x,E]=r.useState(!1),[y,k]=r.useState([]),[j,S]=r.useState([]),O=G(i=>A.put("/api/v1/messages",i),{onSuccess:()=>{alert("Message posted successfully")}});return r.useEffect(()=>{console.log(t),t&&(o(t.title),E(t.published),d(t.series),c(t.preacher),u(t.thumbnail),m(t.date),k([...(t==null?void 0:t.passages)??[]]))},[t]),r.useEffect(()=>{console.log(j)},[j]),e.jsxs(C,{show:s,size:"md",popup:!0,onClose:()=>{k([...(t==null?void 0:t.passages)??[]]),a(!1)},children:[e.jsx(C.Header,{}),e.jsx(C.Body,{children:e.jsxs("div",{className:"space-y-6",children:[e.jsx("h3",{className:"text-xl font-medium text-gray-900 dark:text-white",children:"Edit Message"}),e.jsx("div",{children:e.jsx($e,{checked:x,label:"Published",onChange:E})}),e.jsx(N,{value:n,title:"Title",placeholder:"Enter the title of the message",id:"title",onChange:i=>{o(i.target.value)}}),e.jsx(N,{value:l,title:"Series",placeholder:"Enter the series of the message",id:"series",onChange:i=>{d(i.target.value)}}),e.jsx(N,{value:p,title:"Preacher",placeholder:"Enter the name of the preacher",id:"preacher",onChange:i=>{c(i.target.value)}}),e.jsx(N,{value:v,title:"Date",placeholder:"Enter the display date of the message",id:"date",onChange:i=>{m(i.target.value)}}),e.jsx(N,{value:f,title:"Image",placeholder:"Enter the image cover url of the message",id:"image",onChange:i=>{u(i.target.value)}}),y.map((i,T)=>((i==null?void 0:i.header)??"").length>0?e.jsx(e.Fragment,{children:e.jsx(L,{value:i.header,title:"Header",placeholder:"Enter the header of the passage",id:`edit-header${T+1}`,onChange:F=>{let g={id:i.id,header:F.target.value,type:i.type,order:i.order};y[T]=g,j.find(w=>w.id===g.id)===void 0?S([...j,g]):S([...j.filter(w=>w.id!==g.id),g]),k([...y])}})}):e.jsxs(e.Fragment,{children:[e.jsx(N,{value:i.verse,title:"Verse",placeholder:"Enter the verse of the passage",id:`edit-verse${T+1}`,onChange:F=>{let g={id:i.id,verse:F.target.value,message:i.message,type:i.type,order:i.order};y[T]=g,j.filter(w=>w.id===g.id).length===0?S([...j,g]):S([...j.filter(w=>w.id!==g.id),g]),k([...y])}}),e.jsx(L,{value:i.message,title:"Message",placeholder:"Enter the message of the passage",id:`edit-message${T+1}`,onChange:F=>{let g={id:i.id,verse:i.verse,message:F.target.value,type:i.type,order:i.order};y[T]=g,j.filter(w=>w.id===g.id).length===0?S([...j,g]):S([...j.filter(w=>w.id!==g.id),g]),k([...y])}})]}))]})}),e.jsx(C.Footer,{children:e.jsx(I,{className:"bg-momentum-orange hover:bg-momentum-orange enabled:hover:bg-momentum-orange",onClick:()=>{O.mutate({id:(t==null?void 0:t.id)??"",title:n,series:l,preacher:p,thumbnail:f,date:v,passages:j,published:x}),a(!1)},children:"Update"})})]})}const dt=()=>{const{data:t,isLoading:s}=Z({queryKey:["message"],queryFn:async()=>{const{data:u}=await A.get("/api/v1/messages/admin");return u}}),[a,n]=r.useState(!1),[o,l]=r.useState(!1),[d,p]=r.useState(!1),[c,f]=r.useState(null);return s?e.jsx("div",{children:"Loading..."}):e.jsxs("main",{className:"w-full h-screen",children:[e.jsx(P,{onSigninClick:()=>n(!0)}),e.jsx("div",{className:"w-full flex flex-col justify-normal items-end",children:e.jsx("button",{className:"bg-momentum-orange hover:bg-momentum-orange text-white font-bold py-2 px-4 rounded max-h-10",onClick:()=>{b.currentUser?l(!0):n(!0)},children:"Add Message"})}),e.jsx("div",{className:"w-full h-screen flex flex-col justify-normal items-center",children:t.data.map(u=>e.jsx("div",{children:e.jsx(ct,{message:u,onEditMessage:v=>{b.currentUser?(f(v),p(!0)):n(!0)}})},u.id))}),e.jsx(_,{openModal:a,setOpenModal:n}),e.jsx(lt,{openModal:o,setOpenModal:l}),e.jsx(it,{openModal:d,setOpenModal:p,message:c})]})};function ct({message:t,onEditMessage:s}){return e.jsx(te,{className:"max-w-lg mt-6",renderImage:()=>e.jsx("img",{crossOrigin:"anonymous",src:t.thumbnail,title:"source: imgur.com"}),horizontal:!1,children:e.jsxs("div",{className:"min-w-96 flex justify-between items-center",children:[e.jsxs("div",{children:[e.jsx("h5",{className:"text-2xl font-bold tracking-tight text-gray-900 dark:text-white",children:t.preacher}),e.jsx("p",{className:"font-normal text-gray-700 dark:text-gray-400",children:t.date}),e.jsx("p",{className:"font-normal text-gray-700 dark:text-gray-400",children:t.series}),e.jsx("p",{className:"font-normal text-gray-700 dark:text-gray-400",children:t.title})]}),e.jsx("button",{onClick:()=>{s(t)},className:"bg-momentum-orange hover:bg-momentum-orange text-white font-bold py-2 px-4 rounded max-h-10",children:"Modify"})]})})}function mt({setOpenAuthModal:t}){const[s,a]=r.useState(""),[n,o]=r.useState(""),[l,d]=r.useState("MomentumUsers"),p=G(c=>A.post("/api/v1/notifications",c),{onSuccess:()=>{alert("Notification sent successfully")}});return e.jsx(te,{className:"min-w-[23%]",children:e.jsxs("form",{className:"flex flex-col gap-4",children:[e.jsx("h5",{className:"mb-2 text-3xl font-bold text-gray-900 dark:text-white",children:"Reach all users anytime"}),e.jsx("p",{className:"mb-5 text-base text-gray-500 dark:text-gray-400 sm:text-lg",children:"Notify users on iOS & Android on upcoming events and lastest updates."}),e.jsxs("div",{children:[e.jsx("div",{className:"mb-2 block",children:e.jsx(R,{htmlFor:"title",value:"Title"})}),e.jsx(U,{id:"title",type:"text",required:!0,onChange:c=>{a(c.target.value)}})]}),e.jsxs("div",{children:[e.jsx("div",{className:"mb-2 block",children:e.jsx(R,{htmlFor:"body",value:"Body"})}),e.jsx(ee,{id:"body",rows:5,required:!0,onChange:c=>{o(c.target.value)}})]}),e.jsx("div",{children:e.jsxs("fieldset",{className:"flex max-w-md flex-col gap-4",onChange:c=>{d(c.target.value)},children:[e.jsx("legend",{className:"mb-4",children:"Choose your target audience"}),e.jsxs("div",{className:"flex items-center gap-2",children:[e.jsx(V,{id:"all-users",name:"audience",value:"MomentumUsers",defaultChecked:!0}),e.jsx(R,{htmlFor:"all-users",children:"All Users"})]}),e.jsxs("div",{className:"flex items-center gap-2",children:[e.jsx(V,{id:"android",name:"audience",value:"MomentumAndroid"}),e.jsx(R,{htmlFor:"android",children:"Android"})]}),e.jsxs("div",{className:"flex items-center gap-2",children:[e.jsx(V,{id:"ios",name:"audience",value:"MomentumIOS"}),e.jsx(R,{htmlFor:"ios",children:"iOS"})]})]})}),e.jsx(I,{className:"bg-momentum-orange hover:bg-momentum-orange  enabled:hover:bg-momentum-orange",type:"submit",onClick:c=>{if(c.preventDefault(),b.currentUser){if(s===""||n===""){alert("Please fill in all fields");return}p.mutate({title:s,body:n,topic:l})}else t()},children:"Send Notification"})]})})}const ut=()=>{const[t,s]=r.useState(!1);return e.jsxs("main",{className:"w-full h-screen",children:[e.jsx(P,{onSigninClick:()=>{s(!0)}}),e.jsx("div",{className:"w-full h-screen flex justify-center items-center",children:e.jsx(mt,{setOpenAuthModal:()=>{s(!0)}})}),e.jsx(_,{openModal:t,setOpenModal:s})]})};function h(...t){return Ge(Ye(t))}const ht=Qe("inline-flex items-center justify-center whitespace-nowrap rounded-md text-sm font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50",{variants:{variant:{default:"bg-primary text-primary-foreground hover:bg-primary/90",destructive:"bg-destructive text-destructive-foreground hover:bg-destructive/90",outline:"border border-input bg-background hover:bg-accent hover:text-accent-foreground",secondary:"bg-secondary text-secondary-foreground hover:bg-secondary/80",ghost:"hover:bg-accent hover:text-accent-foreground",link:"text-primary underline-offset-4 hover:underline"},size:{default:"h-10 px-4 py-2",sm:"h-9 rounded-md px-3",lg:"h-11 rounded-md px-8",icon:"h-10 w-10"}},defaultVariants:{variant:"default",size:"default"}}),$=r.forwardRef(({className:t,variant:s,size:a,asChild:n=!1,...o},l)=>{const d=n?Ve:"button";return e.jsx(d,{className:h(ht({variant:s,size:a,className:t})),ref:l,...o})});$.displayName="Button";const he=He,fe=qe,ft=r.forwardRef(({className:t,inset:s,children:a,...n},o)=>e.jsxs(se,{ref:o,className:h("flex cursor-default select-none items-center rounded-sm px-2 py-1.5 text-sm outline-none focus:bg-accent data-[state=open]:bg-accent",s&&"pl-8",t),...n,children:[a,e.jsx(_e,{className:"ml-auto h-4 w-4"})]}));ft.displayName=se.displayName;const xt=r.forwardRef(({className:t,...s},a)=>e.jsx(ae,{ref:a,className:h("z-50 min-w-[8rem] overflow-hidden rounded-md border bg-popover p-1 text-popover-foreground shadow-lg data-[state=open]:animate-in data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=open]:fade-in-0 data-[state=closed]:zoom-out-95 data-[state=open]:zoom-in-95 data-[side=bottom]:slide-in-from-top-2 data-[side=left]:slide-in-from-right-2 data-[side=right]:slide-in-from-left-2 data-[side=top]:slide-in-from-bottom-2",t),...s}));xt.displayName=ae.displayName;const W=r.forwardRef(({className:t,sideOffset:s=4,...a},n)=>e.jsx(Be,{children:e.jsx(re,{ref:n,sideOffset:s,className:h("z-50 min-w-[8rem] overflow-hidden rounded-md border bg-popover p-1 text-popover-foreground shadow-md data-[state=open]:animate-in data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=open]:fade-in-0 data-[state=closed]:zoom-out-95 data-[state=open]:zoom-in-95 data-[side=bottom]:slide-in-from-top-2 data-[side=left]:slide-in-from-right-2 data-[side=right]:slide-in-from-left-2 data-[side=top]:slide-in-from-bottom-2",t),...a})}));W.displayName=re.displayName;const D=r.forwardRef(({className:t,inset:s,...a},n)=>e.jsx(ne,{ref:n,className:h("relative flex cursor-default select-none items-center rounded-sm px-2 py-1.5 text-sm outline-none transition-colors focus:bg-accent focus:text-accent-foreground data-[disabled]:pointer-events-none data-[disabled]:opacity-50",s&&"pl-8",t),...a}));D.displayName=ne.displayName;const xe=r.forwardRef(({className:t,children:s,checked:a,...n},o)=>e.jsxs(oe,{ref:o,className:h("relative flex cursor-default select-none items-center rounded-sm py-1.5 pl-8 pr-2 text-sm outline-none transition-colors focus:bg-accent focus:text-accent-foreground data-[disabled]:pointer-events-none data-[disabled]:opacity-50",t),checked:a,...n,children:[e.jsx("span",{className:"absolute left-2 flex h-3.5 w-3.5 items-center justify-center",children:e.jsx(le,{children:e.jsx(ue,{className:"h-4 w-4"})})}),s]}));xe.displayName=oe.displayName;const pt=r.forwardRef(({className:t,children:s,...a},n)=>e.jsxs(ie,{ref:n,className:h("relative flex cursor-default select-none items-center rounded-sm py-1.5 pl-8 pr-2 text-sm outline-none transition-colors focus:bg-accent focus:text-accent-foreground data-[disabled]:pointer-events-none data-[disabled]:opacity-50",t),...a,children:[e.jsx("span",{className:"absolute left-2 flex h-3.5 w-3.5 items-center justify-center",children:e.jsx(le,{children:e.jsx(We,{className:"h-2 w-2 fill-current"})})}),s]}));pt.displayName=ie.displayName;const pe=r.forwardRef(({className:t,inset:s,...a},n)=>e.jsx(de,{ref:n,className:h("px-2 py-1.5 text-sm font-semibold",s&&"pl-8",t),...a}));pe.displayName=de.displayName;const ge=r.forwardRef(({className:t,...s},a)=>e.jsx(ce,{ref:a,className:h("-mx-1 my-1 h-px bg-muted",t),...s}));ge.displayName=ce.displayName;const be=r.forwardRef(({className:t,type:s,...a},n)=>e.jsx("input",{type:s,className:h("flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50",t),ref:n,...a}));be.displayName="Input";const je=r.forwardRef(({className:t,...s},a)=>e.jsx("div",{className:"relative w-full overflow-auto",children:e.jsx("table",{ref:a,className:h("w-full caption-bottom text-sm",t),...s})}));je.displayName="Table";const ve=r.forwardRef(({className:t,...s},a)=>e.jsx("thead",{ref:a,className:h("[&_tr]:border-b",t),...s}));ve.displayName="TableHeader";const ye=r.forwardRef(({className:t,...s},a)=>e.jsx("tbody",{ref:a,className:h("[&_tr:last-child]:border-0",t),...s}));ye.displayName="TableBody";const gt=r.forwardRef(({className:t,...s},a)=>e.jsx("tfoot",{ref:a,className:h("border-t bg-muted/50 font-medium [&>tr]:last:border-b-0",t),...s}));gt.displayName="TableFooter";const z=r.forwardRef(({className:t,...s},a)=>e.jsx("tr",{ref:a,className:h("border-b transition-colors hover:bg-muted/50 data-[state=selected]:bg-muted",t),...s}));z.displayName="TableRow";const Ne=r.forwardRef(({className:t,...s},a)=>e.jsx("th",{ref:a,className:h("h-12 px-4 text-left align-middle font-medium text-muted-foreground [&:has([role=checkbox])]:pr-0",t),...s}));Ne.displayName="TableHead";const Q=r.forwardRef(({className:t,...s},a)=>e.jsx("td",{ref:a,className:h("p-4 align-middle [&:has([role=checkbox])]:pr-0",t),...s}));Q.displayName="TableCell";const bt=r.forwardRef(({className:t,...s},a)=>e.jsx("caption",{ref:a,className:h("mt-4 text-sm text-muted-foreground",t),...s}));bt.displayName="TableCaption";const Y=r.forwardRef(({className:t,...s},a)=>e.jsx(me,{ref:a,className:h("peer h-4 w-4 shrink-0 rounded-sm border border-primary ring-offset-background focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50 data-[state=checked]:bg-primary data-[state=checked]:text-primary-foreground",t),...s,children:e.jsx(Ke,{className:h("flex items-center justify-center text-current"),children:e.jsx(ue,{className:"h-4 w-4"})})}));Y.displayName=me.displayName;const X=[{id:"select",header:({table:t})=>e.jsx(Y,{checked:t.getIsAllPageRowsSelected()||t.getIsSomePageRowsSelected()&&"indeterminate",onCheckedChange:s=>t.toggleAllPageRowsSelected(!!s),"aria-label":"Select all"}),cell:({row:t})=>e.jsx(Y,{checked:t.getIsSelected(),onCheckedChange:s=>t.toggleSelected(!!s),"aria-label":"Select row"}),enableSorting:!1,enableHiding:!1},{accessorKey:"date",header:"Date",cell:({row:t})=>e.jsx("div",{className:"capitalize",children:b.currentUser?t.getValue("date"):t.getValue("date").split("").map(()=>"*").join("")})},{accessorKey:"user",header:({column:t})=>e.jsxs($,{variant:"ghost",onClick:()=>t.toggleSorting(t.getIsSorted()==="asc"),children:["Fullname",e.jsx(B,{className:"ml-2 h-4 w-4"})]}),cell:({row:t})=>e.jsx("div",{className:"capitalize",children:b.currentUser?t.getValue("user").fullname:t.index==0?"🔒 Signin to view information":""})},{accessorKey:"description",header:"Description",cell:({row:t})=>e.jsx("div",{className:"capitalize",children:b.currentUser?t.getValue("description"):t.getValue("description").split("").map(()=>"*").join("")})},{accessorKey:"user",header:({column:t})=>e.jsxs($,{variant:"ghost",onClick:()=>t.toggleSorting(t.getIsSorted()==="asc"),children:["Email",e.jsx(B,{className:"ml-2 h-4 w-4"})]}),cell:({row:t})=>e.jsx("div",{className:b.currentUser?"lowercase":"capitalize",children:b.currentUser?t.getValue("user").email:t.index==0?"🔒 Signin to view information":""})},{accessorKey:"createdOn",header:({column:t})=>e.jsxs($,{variant:"ghost",onClick:()=>t.toggleSorting(t.getIsSorted()==="asc"),children:["Created On",e.jsx(B,{className:"ml-2 h-4 w-4"})]}),cell:({row:t})=>e.jsx("div",{className:"lowercase",children:b.currentUser?new Date(t.getValue("createdOn")).toLocaleDateString():t.getValue("createdOn").split("").map(()=>"*").join("")})},{accessorKey:"amount",header:()=>e.jsx("div",{className:"text-right",children:"Amount"}),cell:({row:t})=>{const s=parseFloat(t.getValue("amount")),a=new Intl.NumberFormat("en-US",{style:"currency",currency:"USD"}).format(s);return e.jsx("div",{className:"text-right font-medium",children:b.currentUser?a:a.split("").map(()=>"$").slice(0,a.length-2).join("")})}},{id:"actions",enableHiding:!1,cell:({row:t})=>{const s=t.original;return e.jsxs(he,{children:[e.jsx(fe,{asChild:!0,children:e.jsxs($,{variant:"ghost",className:"h-8 w-8 p-0",children:[e.jsx("span",{className:"sr-only",children:"Open menu"}),e.jsx(Je,{className:"h-4 w-4"})]})}),e.jsxs(W,{align:"end",children:[e.jsx(pe,{children:"Actions"}),e.jsx(D,{onClick:()=>navigator.clipboard.writeText(s.id),children:"Copy payment ID"}),e.jsx(ge,{}),e.jsx(D,{children:"View customer"}),e.jsx(D,{children:"View payment details"})]})]})}}];function jt({data:t}){var u,v;const[s,a]=r.useState([]),[n,o]=r.useState([]),[l,d]=r.useState({}),[p,c]=r.useState({}),f=Oe({data:t,columns:X,onSortingChange:a,onColumnFiltersChange:o,getCoreRowModel:De(),getPaginationRowModel:ze(),getSortedRowModel:Ue(),getFilteredRowModel:Le(),onColumnVisibilityChange:d,onRowSelectionChange:c,state:{sorting:s,columnFilters:n,columnVisibility:l,rowSelection:p}});return e.jsxs("div",{className:"w-full",children:[e.jsxs("div",{className:"flex items-center py-4",children:[e.jsx(be,{placeholder:"Filter emails...",value:((u=f.getColumn("email"))==null?void 0:u.getFilterValue())??"",onChange:m=>{var x;return(x=f.getColumn("email"))==null?void 0:x.setFilterValue(m.target.value)},className:"max-w-sm"}),e.jsxs(he,{children:[e.jsx(fe,{asChild:!0,children:e.jsxs($,{variant:"outline",className:"ml-auto",children:["Columns ",e.jsx(Xe,{className:"ml-2 h-4 w-4"})]})}),e.jsx(W,{align:"end",children:f.getAllColumns().filter(m=>m.getCanHide()).map(m=>e.jsx(xe,{className:"capitalize",checked:m.getIsVisible(),onCheckedChange:x=>m.toggleVisibility(!!x),children:m.id},m.id))})]})]}),e.jsx("div",{className:"rounded-md border",children:e.jsxs(je,{children:[e.jsx(ve,{children:f.getHeaderGroups().map(m=>e.jsx(z,{children:m.headers.map(x=>e.jsx(Ne,{children:x.isPlaceholder?null:J(x.column.columnDef.header,x.getContext())},x.id))},m.id))}),e.jsx(ye,{children:(v=f.getRowModel().rows)!=null&&v.length?f.getRowModel().rows.map(m=>e.jsx(z,{"data-state":m.getIsSelected()&&"selected",children:m.getVisibleCells().map(x=>e.jsx(Q,{children:J(x.column.columnDef.cell,x.getContext())},x.id))},m.id)):e.jsx(z,{children:e.jsx(Q,{colSpan:X.length,className:"h-24 text-center",children:"No results."})})})]})}),e.jsxs("div",{className:"flex items-center justify-end space-x-2 py-4",children:[e.jsxs("div",{className:"flex-1 text-sm text-muted-foreground",children:[f.getFilteredSelectedRowModel().rows.length," of"," ",f.getFilteredRowModel().rows.length," row(s) selected."]}),e.jsxs("div",{className:"space-x-2",children:[e.jsx($,{variant:"outline",size:"sm",onClick:()=>f.previousPage(),disabled:!f.getCanPreviousPage(),children:"Previous"}),e.jsx($,{variant:"outline",size:"sm",onClick:()=>f.nextPage(),disabled:!f.getCanNextPage(),children:"Next"})]})]})]})}const vt=()=>{const{data:t,isLoading:s}=Z({queryKey:["transactions"],queryFn:async()=>{const{data:o}=await A.get("/api/v1/transactions");return o}}),[a,n]=r.useState(!1);return s?e.jsx("div",{className:"w-full h-screen",children:"Loading..."}):e.jsxs("main",{className:"w-full h-screen",children:[e.jsx(P,{onSigninClick:()=>{n(!0)}}),e.jsx(jt,{data:t}),e.jsx(_,{openModal:a,setOpenModal:n})]})},yt=()=>e.jsx("div",{className:"w-full",children:e.jsx(P,{onSigninClick:()=>{}})}),Nt=ke([{path:q,element:e.jsx(dt,{})},{path:st,element:e.jsx(ot,{})},{path:H,element:e.jsx(ut,{})},{path:K,element:e.jsx(vt,{})},{path:tt,element:e.jsx(yt,{})}]);A.interceptors.request.use(t=>(t.headers.Authorization=`Bearer ${et}`,t),t=>Promise.reject(t));function wt(){return e.jsxs(Me,{client:Ze,children:[e.jsx(Pe,{router:Nt}),e.jsx(Re.ReactQueryDevtools,{initialIsOpen:!1})]})}Ce.createRoot(document.getElementById("root")).render(e.jsx(we.StrictMode,{children:e.jsx(wt,{})}));
