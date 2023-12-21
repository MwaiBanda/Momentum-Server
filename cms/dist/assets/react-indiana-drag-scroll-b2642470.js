import{a as f,r as g}from"./react-bc4662ae.js";/*! *****************************************************************************
Copyright (c) Microsoft Corporation.

Permission to use, copy, modify, and/or distribute this software for any
purpose with or without fee is hereby granted.

THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES WITH
REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF MERCHANTABILITY
AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY SPECIAL, DIRECT,
INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES WHATSOEVER RESULTING FROM
LOSS OF USE, DATA OR PROFITS, WHETHER IN AN ACTION OF CONTRACT, NEGLIGENCE OR
OTHER TORTIOUS ACTION, ARISING OUT OF OR IN CONNECTION WITH THE USE OR
PERFORMANCE OF THIS SOFTWARE.
***************************************************************************** */var v=function(s,r){return(v=Object.setPrototypeOf||{__proto__:[]}instanceof Array&&function(n,t){n.__proto__=t}||function(n,t){for(var o in t)t.hasOwnProperty(o)&&(n[o]=t[o])})(s,r)},m,u,b=(function(s){/*!
  Copyright (c) 2017 Jed Watson.
  Licensed under the MIT License (MIT), see
  http://jedwatson.github.io/classnames
*/(function(){var r={}.hasOwnProperty;function n(){for(var t=[],o=0;o<arguments.length;o++){var e=arguments[o];if(e){var i=typeof e;if(i==="string"||i==="number")t.push(e);else if(Array.isArray(e)&&e.length){var a=n.apply(null,e);a&&t.push(a)}else if(i==="object")for(var l in e)r.call(e,l)&&e[l]&&t.push(l)}}return t.join(" ")}s.exports?(n.default=n,s.exports=n):window.classNames=n})()}(u={path:m,exports:{},require:function(s,r){return function(){throw new Error("Dynamic requires are not currently supported by @rollup/plugin-commonjs")}(r==null&&u.path)}},u.exports),u.exports);function h(s,r,n){var t,o,e,i,a;function l(){var c=Date.now()-i;c<r&&c>=0?t=setTimeout(l,r-c):(t=null,n||(a=s.apply(e,o),e=o=null))}r==null&&(r=100);var p=function(){e=this,o=arguments,i=Date.now();var c=n&&!t;return t||(t=setTimeout(l,r)),c&&(a=s.apply(e,o),e=o=null),a};return p.clear=function(){t&&(clearTimeout(t),t=null)},p.flush=function(){t&&(a=s.apply(e,o),e=o=null,clearTimeout(t),t=null)},p}h.debounce=h;var y=h;(function(s,r){r===void 0&&(r={});var n=r.insertAt;if(s&&typeof document<"u"){var t=document.head||document.getElementsByTagName("head")[0],o=document.createElement("style");o.type="text/css",n==="top"&&t.firstChild?t.insertBefore(o,t.firstChild):t.appendChild(o),o.styleSheet?o.styleSheet.cssText=s:o.appendChild(document.createTextNode(s))}})(`.indiana-scroll-container {
  overflow: auto; }
  .indiana-scroll-container--dragging {
    scroll-behavior: auto !important; }
    .indiana-scroll-container--dragging > * {
      pointer-events: none;
      cursor: -webkit-grab;
      cursor: grab; }
  .indiana-scroll-container--hide-scrollbars {
    overflow: hidden;
    overflow: -moz-scrollbars-none;
    -ms-overflow-style: none;
    scrollbar-width: none; }
    .indiana-scroll-container--hide-scrollbars::-webkit-scrollbar {
      display: none !important;
      height: 0 !important;
      width: 0 !important;
      background: transparent !important;
      -webkit-appearance: none !important; }
  .indiana-scroll-container--native-scroll {
    overflow: auto; }

.indiana-dragging {
  cursor: -webkit-grab;
  cursor: grab; }
`);var d,w=(d="indiana-scroll-container",function(s,r){if(!s)return d;var n;typeof s=="string"?n=s:r=s;var t=d;return n&&(t+="__"+n),t+(r?Object.keys(r).reduce(function(o,e){var i=r[e];return i&&(o+=" "+(typeof i=="boolean"?t+"--"+e:t+"--"+e+"_"+i)),o},""):"")});(function(s){function r(n){var t=s.call(this,n)||this;return t.onEndScroll=function(){t.scrolling=!1,!t.pressed&&t.started&&t.processEnd()},t.onScroll=function(o){var e=t.container.current;e.scrollLeft===t.scrollLeft&&e.scrollTop===t.scrollTop||(t.scrolling=!0,t.processScroll(o),t.onEndScroll())},t.onTouchStart=function(o){var e=t.props.nativeMobileScroll;if(t.isDraggable(o.target))if(t.internal=!0,e&&t.scrolling)t.pressed=!0;else{var i=o.touches[0];t.processClick(o,i.clientX,i.clientY),!e&&t.props.stopPropagation&&o.stopPropagation()}},t.onTouchEnd=function(o){var e=t.props.nativeMobileScroll;t.pressed&&(!t.started||t.scrolling&&e?t.pressed=!1:t.processEnd(),t.forceUpdate())},t.onTouchMove=function(o){var e=t.props.nativeMobileScroll;if(t.pressed&&(!e||!t.isMobile)){var i=o.touches[0];i&&t.processMove(o,i.clientX,i.clientY),o.preventDefault(),t.props.stopPropagation&&o.stopPropagation()}},t.onMouseDown=function(o){t.isDraggable(o.target)&&t.isScrollable()&&(t.internal=!0,t.props.buttons.indexOf(o.button)!==-1&&(t.processClick(o,o.clientX,o.clientY),o.preventDefault(),t.props.stopPropagation&&o.stopPropagation()))},t.onMouseMove=function(o){t.pressed&&(t.processMove(o,o.clientX,o.clientY),o.preventDefault(),t.props.stopPropagation&&o.stopPropagation())},t.onMouseUp=function(o){t.pressed&&(t.started?t.processEnd():(t.internal=!1,t.pressed=!1,t.forceUpdate(),t.props.onClick&&t.props.onClick(o)),o.preventDefault(),t.props.stopPropagation&&o.stopPropagation())},t.container=f.createRef(),t.onEndScroll=y(t.onEndScroll,300),t.scrolling=!1,t.started=!1,t.pressed=!1,t.internal=!1,t.getRef=t.getRef.bind(t),t}return function(n,t){function o(){this.constructor=n}v(n,t),n.prototype=t===null?Object.create(t):(o.prototype=t.prototype,new o)}(r,s),r.prototype.componentDidMount=function(){var n=this.props.nativeMobileScroll,t=this.container.current;window.addEventListener("mouseup",this.onMouseUp),window.addEventListener("mousemove",this.onMouseMove),window.addEventListener("touchmove",this.onTouchMove,{passive:!1}),window.addEventListener("touchend",this.onTouchEnd),t.addEventListener("touchstart",this.onTouchStart,{passive:!1}),t.addEventListener("mousedown",this.onMouseDown,{passive:!1}),n&&(this.isMobile=this.isMobileDevice(),this.isMobile&&this.forceUpdate())},r.prototype.componentWillUnmount=function(){window.removeEventListener("mouseup",this.onMouseUp),window.removeEventListener("mousemove",this.onMouseMove),window.removeEventListener("touchmove",this.onTouchMove),window.removeEventListener("touchend",this.onTouchEnd)},r.prototype.getElement=function(){return this.container.current},r.prototype.isMobileDevice=function(){return window.orientation!==void 0||navigator.userAgent.indexOf("IEMobile")!==-1},r.prototype.isDraggable=function(n){var t=this.props.ignoreElements;if(t){var o=n.closest(t);return o===null||o.contains(this.getElement())}return!0},r.prototype.isScrollable=function(){var n=this.container.current;return n&&(n.scrollWidth>n.clientWidth||n.scrollHeight>n.clientHeight)},r.prototype.processClick=function(n,t,o){var e=this.container.current;this.scrollLeft=e.scrollLeft,this.scrollTop=e.scrollTop,this.clientX=t,this.clientY=o,this.pressed=!0},r.prototype.processStart=function(n){n===void 0&&(n=!0);var t=this.props.onStartScroll;this.started=!0,n&&document.body.classList.add("indiana-dragging"),t&&t({external:!this.internal}),this.forceUpdate()},r.prototype.processScroll=function(n){if(this.started){var t=this.props.onScroll;t&&t({external:!this.internal})}else this.processStart(!1)},r.prototype.processMove=function(n,t,o){var e=this.props,i=e.horizontal,a=e.vertical,l=e.activationDistance,p=e.onScroll,c=this.container.current;this.started?(i&&(c.scrollLeft-=t-this.clientX),a&&(c.scrollTop-=o-this.clientY),p&&p({external:!this.internal}),this.clientX=t,this.clientY=o,this.scrollLeft=c.scrollLeft,this.scrollTop=c.scrollTop):(i&&Math.abs(t-this.clientX)>l||a&&Math.abs(o-this.clientY)>l)&&(this.clientX=t,this.clientY=o,this.processStart())},r.prototype.processEnd=function(){var n=this.props.onEndScroll;this.container.current&&n&&n({external:!this.internal}),this.pressed=!1,this.started=!1,this.scrolling=!1,this.internal=!1,document.body.classList.remove("indiana-dragging"),this.forceUpdate()},r.prototype.getRef=function(n){[this.container,this.props.innerRef].forEach(function(t){t&&(typeof t=="function"?t(n):t.current=n)})},r.prototype.render=function(){var n=this.props,t=n.children,o=n.draggingClassName,e=n.className,i=n.style,a=n.hideScrollbars,l=n.component;return f.createElement(l,{className:b(e,this.pressed&&o,w({dragging:this.pressed,"hide-scrollbars":a,"native-scroll":this.isMobile})),style:i,ref:this.getRef,onScroll:this.onScroll},t)},r.defaultProps={nativeMobileScroll:!0,hideScrollbars:!0,activationDistance:10,vertical:!0,horizontal:!0,stopPropagation:!1,style:{},component:"div",buttons:[0]},r})(g.PureComponent);
