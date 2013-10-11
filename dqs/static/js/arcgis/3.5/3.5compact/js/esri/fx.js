/*
 COPYRIGHT 2009 ESRI

 TRADE SECRETS: ESRI PROPRIETARY AND CONFIDENTIAL
 Unpublished material - all rights reserved under the
 Copyright Laws of the United States and applicable international
 laws, treaties, and conventions.

 For additional information, contact:
 Environmental Systems Research Institute, Inc.
 Attn: Contracts and Legal Services Department
 380 New York Street
 Redlands, California, 92373
 USA

 email: contracts@esri.com
 */
//>>built
define("esri/fx",["dojo/_base/connect","dojo/_base/fx","dojo/_base/lang","dojo/dom","dojo/dom-geometry","dojo/dom-style","dojo/fx","dojo/has","esri/kernel"],function(_1,_2,_3,_4,_5,_6,_7,_8,_9){var fx={animateRange:function(_a){var _b=_a.range;return new _2.Animation(_3.mixin({curve:new _2._Line(_b.start,_b.end)},_a));},resize:function(_c){var _d=(_c.node=_4.byId(_c.node)),_e=_c.start,_f=_c.end;if(!_e){var mb=_5.getMarginBox(_d),pb=_5.getPadBorderExtents(_d);_e=(_c.start={left:mb.l+pb.l,top:mb.t+pb.t,width:mb.w-pb.w,height:mb.h-pb.h});}if(!_f){var _10=_c.anchor?_c.anchor:{x:_e.left,y:_e.top},_11=_c.size;_f=_c.end={left:(_e.left-((_11.width-_e.width)*(_10.x-_e.left)/_e.width)),top:(_e.top-((_11.height-_e.height)*(_10.y-_e.top)/_e.height)),width:_11.width,height:_11.height};}return _2.animateProperty(_3.mixin({properties:{left:{start:_e.left,end:_f.left},top:{start:_e.top,end:_f.top},width:{start:_e.width,end:_f.width},height:{start:_e.height,end:_f.height}}},_c));},slideTo:function(_12){var _13=(_12.node=_4.byId(_12.node)),_14=_6.getComputedStyle,top=null,_15=null,_16=(function(){var _17=_13;return function(){var pos=_17.style.position=="absolute"?"absolute":"relative";top=(pos=="absolute"?_13.offsetTop:parseInt(_14(_13).top)||0);_15=(pos=="absolute"?_13.offsetLeft:parseInt(_14(_13).left)||0);if(pos!="absolute"&&pos!="relative"){var ret=_5.position(_17,true);top=ret.y;_15=ret.x;_17.style.position="absolute";_17.style.top=top+"px";_17.style.left=_15+"px";}};}());_16();var _18=_2.animateProperty(_3.mixin({properties:{top:{start:top,end:_12.top||0},left:{start:_15,end:_12.left||0}}},_12));_1.connect(_18,"beforeBegin",_18,_16);return _18;},flash:function(_19){_19=_3.mixin({end:"#f00",duration:500,count:1},_19);_19.duration/=_19.count*2;var _1a=_4.byId(_19.node),_1b=_19.start;if(!_1b){_1b=_6.getComputedStyle(_1a).backgroundColor;}var end=_19.end,_1c=_19.duration,_1d=[],i,il=_19.count,_1e={node:_1a,duration:_1c};for(i=0;i<il;i++){_1d.push(_2.animateProperty(_3.mixin({properties:{backgroundColor:{start:_1b,end:end}}},_1e)));_1d.push(_2.animateProperty(_3.mixin({properties:{backgroundColor:{start:end,end:_1b}}},_1e)));}return _7.chain(_1d);}};if(_8("extend-esri")){_3.mixin(_3.getObject("fx",true,_9),fx);}return fx;});