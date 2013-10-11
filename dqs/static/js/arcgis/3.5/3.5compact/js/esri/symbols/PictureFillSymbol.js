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
define("esri/symbols/PictureFillSymbol",["dojo/_base/declare","dojo/_base/lang","dojo/sniff","dojox/gfx/_base","esri/kernel","esri/lang","esri/urlUtils","esri/symbols/FillSymbol"],function(_1,_2,_3,_4,_5,_6,_7,_8){var _9={xoffset:0,yoffset:0,width:12,height:12};var _a=_1(_8,{declaredClass:"esri.symbol.PictureFillSymbol",type:"picturefillsymbol",xscale:1,yscale:1,xoffset:0,yoffset:0,constructor:function(_b,_c,_d,_e){if(_b){if(_2.isString(_b)){this.url=_b;if(_c!==undefined){this.outline=_c;}if(_d!==undefined){this.width=_d;}if(_e!==undefined){this.height=_e;}}else{this.xoffset=_4.pt2px(_b.xoffset);this.yoffset=_4.pt2px(_b.yoffset);this.width=_4.pt2px(_b.width);this.height=_4.pt2px(_b.height);var _f=_b.imageData;if((!(_3("ie")<9))&&_f){var _10=this.url;this.url="data:"+(_b.contentType||"image")+";base64,"+_f;this.imageData=_10;}}}else{_2.mixin(this,_9);this.width=_4.pt2px(this.width);this.height=_4.pt2px(this.height);}},setWidth:function(_11){this.width=_11;return this;},setHeight:function(_12){this.height=_12;return this;},setOffset:function(x,y){this.xoffset=x;this.yoffset=y;return this;},setUrl:function(url){if(url!==this.url){delete this.imageData;delete this.contentType;}this.url=url;return this;},setXScale:function(_13){this.xscale=_13;return this;},setYScale:function(_14){this.yscale=_14;return this;},getStroke:function(){return this.outline&&this.outline.getStroke();},getFill:function(){return _2.mixin({},_4.defaultPattern,{src:this.url,width:(this.width*this.xscale),height:(this.height*this.yscale),x:this.xoffset,y:this.yoffset});},getShapeDescriptors:function(){return {defaultShape:{type:"path",path:"M -10,-10 L 10,0 L 10,10 L -10,10 L -10,-10 E"},fill:this.getFill(),stroke:this.getStroke()};},toJson:function(){var url=this.url,_15=this.imageData;if(url.indexOf("data:")===0){var _16=url;url=_15;var _17=_16.indexOf(";base64,")+8;_15=_16.substr(_17);}url=_7.getAbsoluteUrl(url);var _18=_4.px2pt(this.width);_18=isNaN(_18)?undefined:_18;var _19=_4.px2pt(this.height);_19=isNaN(_19)?undefined:_19;var _1a=_4.px2pt(this.xoffset);_1a=isNaN(_1a)?undefined:_1a;var _1b=_4.px2pt(this.yoffset);_1b=isNaN(_1b)?undefined:_1b;var _1c=_6.fixJson(_2.mixin(this.inherited("toJson",arguments),{type:"esriPFS",url:url,imageData:_15,contentType:this.contentType,width:_18,height:_19,xoffset:_1a,yoffset:_1b,xscale:this.xscale,yscale:this.yscale}));if(!_1c.imageData){delete _1c.imageData;}return _1c;}});_a.defaultProps=_9;if(_3("extend-esri")){_2.setObject("symbol.PictureFillSymbol",_a,_5);_5.symbol.defaultPictureFillSymbol=_9;}return _a;});