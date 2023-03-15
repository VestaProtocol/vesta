"use strict";(self.webpackChunkvesta_docs=self.webpackChunkvesta_docs||[]).push([[374],{3905:(e,t,n)=>{n.d(t,{Zo:()=>d,kt:()=>m});var r=n(7294);function a(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function o(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function i(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?o(Object(n),!0).forEach((function(t){a(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):o(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function l(e,t){if(null==e)return{};var n,r,a=function(e,t){if(null==e)return{};var n,r,a={},o=Object.keys(e);for(r=0;r<o.length;r++)n=o[r],t.indexOf(n)>=0||(a[n]=e[n]);return a}(e,t);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(r=0;r<o.length;r++)n=o[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(a[n]=e[n])}return a}var s=r.createContext({}),c=function(e){var t=r.useContext(s),n=t;return e&&(n="function"==typeof e?e(t):i(i({},t),e)),n},d=function(e){var t=c(e.components);return r.createElement(s.Provider,{value:t},e.children)},u="mdxType",p={inlineCode:"code",wrapper:function(e){var t=e.children;return r.createElement(r.Fragment,{},t)}},h=r.forwardRef((function(e,t){var n=e.components,a=e.mdxType,o=e.originalType,s=e.parentName,d=l(e,["components","mdxType","originalType","parentName"]),u=c(n),h=a,m=u["".concat(s,".").concat(h)]||u[h]||p[h]||o;return n?r.createElement(m,i(i({ref:t},d),{},{components:n})):r.createElement(m,i({ref:t},d))}));function m(e,t){var n=arguments,a=t&&t.mdxType;if("string"==typeof e||a){var o=n.length,i=new Array(o);i[0]=h;var l={};for(var s in t)hasOwnProperty.call(t,s)&&(l[s]=t[s]);l.originalType=e,l[u]="string"==typeof e?e:a,i[1]=l;for(var c=2;c<o;c++)i[c]=n[c];return r.createElement.apply(null,i)}return r.createElement.apply(null,n)}h.displayName="MDXCreateElement"},9122:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>s,contentTitle:()=>i,default:()=>p,frontMatter:()=>o,metadata:()=>l,toc:()=>c});var r=n(7462),a=(n(7294),n(3905));const o={sidebar_position:4},i="Building an Atomic Swap",l={unversionedId:"contracts/atomicswap",id:"contracts/atomicswap",title:"Building an Atomic Swap",description:"What is an Atomic Swap?",source:"@site/docs/3-contracts/4-atomicswap.md",sourceDirName:"3-contracts",slug:"/contracts/atomicswap",permalink:"/docs/contracts/atomicswap",draft:!1,editUrl:"https://github.com/VestaProtocol/vesta/tree/master/assets/vesta-docs/docs/3-contracts/4-atomicswap.md",tags:[],version:"current",sidebarPosition:4,frontMatter:{sidebar_position:4},sidebar:"tutorialSidebar",previous:{title:"NFT Marketplace",permalink:"/docs/contracts/marketplace"},next:{title:"Setting Up",permalink:"/docs/nodes/install"}},s={},c=[{value:"What is an Atomic Swap?",id:"what-is-an-atomic-swap",level:2},{value:"The Smart Contract",id:"the-smart-contract",level:2},{value:"Defining Data-Structures",id:"defining-data-structures",level:3},{value:"Data Permanence",id:"data-permanence",level:3},{value:"Creating Orders",id:"creating-orders",level:3},{value:"Fulfilling Orders",id:"fulfilling-orders",level:3},{value:"Querying Data",id:"querying-data",level:3},{value:"Interacting With The Contract",id:"interacting-with-the-contract",level:2},{value:"Creating Swap",id:"creating-swap",level:3},{value:"Creating Order",id:"creating-order",level:3},{value:"Fulfilling Order",id:"fulfilling-order",level:3},{value:"Conclusion",id:"conclusion",level:2}],d={toc:c},u="wrapper";function p(e){let{components:t,...n}=e;return(0,a.kt)(u,(0,r.Z)({},d,n,{components:t,mdxType:"MDXLayout"}),(0,a.kt)("h1",{id:"building-an-atomic-swap"},"Building an Atomic Swap"),(0,a.kt)("h2",{id:"what-is-an-atomic-swap"},"What is an Atomic Swap?"),(0,a.kt)("p",null,"An atomic swap is a token swapping mechanism that allows users to create swap orders that can then be fulfilled by\nother users. For example, if I have 10 Atom tokens, and I want 15 Vesta tokens, I can create an order for 15000000uvst\nand 10000000uatom. Then, if someone decides that want 10 Atom tokens and has 15 Vesta tokens, they can fulfil this order\nand the Vesta tokens will be sent to me, and they'll get my Atom tokens."),(0,a.kt)("h2",{id:"the-smart-contract"},"The Smart Contract"),(0,a.kt)("h3",{id:"defining-data-structures"},"Defining Data-Structures"),(0,a.kt)("p",null,"We first need to think about what each order will be composed of. Let's think about this, each order will need a\n",(0,a.kt)("em",{parentName:"p"},"token wanted")," and a ",(0,a.kt)("em",{parentName:"p"},"token offered"),", let's call these ",(0,a.kt)("inlineCode",{parentName:"p"},"tokenIn")," and ",(0,a.kt)("inlineCode",{parentName:"p"},"tokenOut"),". We will also need to keep track of who\nmade the order, let's call this ",(0,a.kt)("inlineCode",{parentName:"p"},"creator"),". And finally, we'll need to keep track of each order with an identifier, let's\nuse the hash of all other details and call it ",(0,a.kt)("inlineCode",{parentName:"p"},"id"),"."),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-js"},"function Order(tokenIn, tokenOut, creator) {\n    let id = STD.crypto.sha256(tokenIn + tokenOut + creator)\n    return {\n        id: id,\n        tokenIn: tokenIn,\n        tokenOut: tokenOut,\n        creator: creator,\n    }\n}\n")),(0,a.kt)("p",null,"The code above will allow use to create an order object using the standard libraries ",(0,a.kt)("inlineCode",{parentName:"p"},"crypto")," library."),(0,a.kt)("h3",{id:"data-permanence"},"Data Permanence"),(0,a.kt)("p",null,"We can't expect each order to happen within the same block as each-other, and as such we'll need a way to save the\norders onto the chain. We can do this with a few helper functions:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-js"},"function SaveOrder(order) {\n    STD.write(order.id, JSON.stringify(order))\n}\n\nfunction RemoveOrder(id) {\n    STD.delete(id)\n}\n\nfunction LoadOrder(id) {\n    let sOrder = STD.read(id)\n    return JSON.parse(sOrder)\n}\n")),(0,a.kt)("h3",{id:"creating-orders"},"Creating Orders"),(0,a.kt)("p",null,"When creating orders, we need to allow users to specify the tokens wanted/offered as well as keep their offer in the\ncontracts account as escrow. We then create the order and save it to the chain, nothing crazy here."),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-js"},'CONTRACT.functions.create = function(tIn, tOut) {\n    let order = Order(tIn, tOut, CTX.sender)\n\n    let ok = STD.bank.sendTokens(CONTRACT.address, tOut)\n    if (!ok) {\n        STD.panic("not enough balance of " + cost)\n    }\n\n    SaveOrder(order)\n}\n')),(0,a.kt)("h3",{id:"fulfilling-orders"},"Fulfilling Orders"),(0,a.kt)("p",null,"In order to fulfil orders, we ask the user for an order ID, then we make sure they have enough tokens to support the\nswap. We then take the tokens from them, send them to the order ",(0,a.kt)("inlineCode",{parentName:"p"},"creator")," and send the tokens they gave the contract\nearlier to the fulfiller. Finally, we delete the order to ensure nobody can try to fulfil it a second time."),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-js"},'CONTRACT.functions.fulfil = function(orderId) {\n    let order = LoadOrder(orderId)\n\n    let ok = STD.bank.sendTokens(CONTRACT.address, order.tokenIn)\n    if (!ok) {\n        STD.panic("not enough balance of " + order.tokenIn)\n    }\n\n    ok = STD.bank.withdrawTokens(order.creator, order.tokenIn)\n    if (!ok) {\n        STD.panic("not enough balance of " + order.tokenIn)\n    }\n\n    ok = STD.bank.withdrawTokens(CTX.sender, order.tokenOut)\n    if (!ok) {\n        STD.panic("not enough balance of " + order.tokenOut)\n    }\n\n    RemoveOrder(orderId)\n}\n')),(0,a.kt)("h3",{id:"querying-data"},"Querying Data"),(0,a.kt)("p",null,"If we want users to be able to view all the info about an order before fulfilling it, we can create a query like this."),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-js"},"CONTRACT.queries.show = function(orderId) {\n    let order = LoadOrder(orderId)\n    return JSON.stringify(order)\n}\n")),(0,a.kt)("h2",{id:"interacting-with-the-contract"},"Interacting With The Contract"),(0,a.kt)("p",null,"Now that you've built an atomic swap, you want to be able to interact with it. To do this you'll need two accounts with different tokens on each of them. (See ",(0,a.kt)("a",{parentName:"p",href:"https://github.com/VestaProtocol/vesta/blob/master/scripts/test_atomswap.sh"},"Testing AtomSwap Script"),")."),(0,a.kt)("h3",{id:"creating-swap"},"Creating Swap"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-shell"},'vestad tx vm store ./examples/orderbook.js --from {account}\nvestad tx vm instantiate book {code} "" --from {account}\n')),(0,a.kt)("h3",{id:"creating-order"},"Creating Order"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-shell"},'vestad tx vm execute book create "{token_a},{token_b}" --from {account}\n')),(0,a.kt)("p",null,"After executing this command you can find the ID by checking this commands output for your contract's storage."),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-shell"},"vestad q tx list-romdata\n")),(0,a.kt)("h3",{id:"fulfilling-order"},"Fulfilling Order"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-shell"},'vestad tx vm execute book fulfil "{id}" --from {account}\n')),(0,a.kt)("h2",{id:"conclusion"},"Conclusion"),(0,a.kt)("p",null,"And there you have it, a full atomic swap system build with JavaScript deployed on Vesta! Happy coding!"))}p.isMDXComponent=!0}}]);