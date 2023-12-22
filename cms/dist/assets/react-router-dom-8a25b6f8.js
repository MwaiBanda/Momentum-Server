import"./react-6a422704.js";import{m as c}from"./react-router-33fe3b2e.js";import{c as f,b as d,E as m}from"./@remix-run-0093eeff.js";/**
 * React Router DOM v6.15.0
 *
 * Copyright (c) Remix Software Inc.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE.md file in the root directory of this source tree.
 *
 * @license MIT
 */function o(){return o=Object.assign?Object.assign.bind():function(r){for(var e=1;e<arguments.length;e++){var n=arguments[e];for(var s in n)Object.prototype.hasOwnProperty.call(n,s)&&(r[s]=n[s])}return r},o.apply(this,arguments)}function v(r,e){return f({basename:e==null?void 0:e.basename,future:o({},e==null?void 0:e.future,{v7_prependBasename:!0}),history:d({window:e==null?void 0:e.window}),hydrationData:(e==null?void 0:e.hydrationData)||p(),routes:r,mapRouteProperties:c}).initialize()}function p(){var r;let e=(r=window)==null?void 0:r.__staticRouterHydrationData;return e&&e.errors&&(e=o({},e,{errors:w(e.errors)})),e}function w(r){if(!r)return null;let e=Object.entries(r),n={};for(let[s,t]of e)if(t&&t.__type==="RouteErrorResponse")n[s]=new m(t.status,t.statusText,t.data,t.internal===!0);else if(t&&t.__type==="Error"){if(t.__subType){let i=window[t.__subType];if(typeof i=="function")try{let a=new i(t.message);a.stack="",n[s]=a}catch{}}if(n[s]==null){let i=new Error(t.message);i.stack="",n[s]=i}}else n[s]=t;return n}var u;(function(r){r.UseScrollRestoration="useScrollRestoration",r.UseSubmit="useSubmit",r.UseSubmitFetcher="useSubmitFetcher",r.UseFetcher="useFetcher"})(u||(u={}));var l;(function(r){r.UseFetchers="useFetchers",r.UseScrollRestoration="useScrollRestoration"})(l||(l={}));export{v as c};
