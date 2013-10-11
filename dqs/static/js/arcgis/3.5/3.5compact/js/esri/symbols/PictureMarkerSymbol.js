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
define("esri/symbols/PictureMarkerSymbol",["dojo/_base/declare","dojo/_base/lang","dojo/sniff","dojox/gfx/_base","esri/kernel","esri/lang","esri/urlUtils","esri/symbols/MarkerSymbol"],function(_1,_2,_3,_4,_5,_6,_7,_8){var _9={url:"",width:12,height:12,angle:0,xoffset:0,yoffset:0};var _a=_1(_8,{declaredClass:"esri.symbol.PictureMarkerSymbol",type:"picturemarkersymbol",constructor:function(_b,_c,_d){if(_b){if(_2.isString(_b)){this.url=_b;if(_c){this.width=_c;}if(_d){this.height=_d;}}else{this.width=_4.pt2px(_b.width);this.height=_4.pt2px(_b.height);var _e=_b.imageData;if((!(_3("ie")<9))&&_e){var _f=this.url;this.url="data:"+(_b.contentType||"image")+";base64,"+_e;this.imageData=_f;}}}else{_2.mixin(this,_9);this.width=_4.pt2px(this.width);this.height=_4.pt2px(this.height);}},getStroke:function(){return null;},getFill:function(){return null;},setWidth:function(_10){this.width=_10;return this;},setHeight:function(_11){this.height=_11;return this;},setUrl:function(url){if(url!==this.url){delete this.imageData;delete this.contentType;}this.url=url;return this;},getShapeDescriptors:function(){var _12={type:"image",x:-Math.round(this.width/2),y:-Math.round(this.height/2),width:this.width,height:this.height,src:this.url||""};return {defaultShape:_12,fill:null,stroke:null};},toJson:function(){var url=this.url,_13=this.imageData;if(url.indexOf("data:")===0){var _14=url;url=_13;var _15=_14.indexOf(";base64,")+8;_13=_14.substr(_15);}url=_7.getAbsoluteUrl(url);var _16=_4.px2pt(this.width);_16=isNaN(_16)?undefined:_16;var _17=_4.px2pt(this.height);_17=isNaN(_17)?undefined:_17;var _18=_6.fixJson(_2.mixin(this.inherited("toJson",arguments),{type:"esriPMS",url:url,imageData:_13,contentType:this.contentType,width:_16,height:_17}));delete _18.color;delete _18.size;if(!_18.imageData){delete _18.imageData;}return _18;}});_a.defaultProps=_9;if(_3("extend-esri")){_2.setObject("symbol.PictureMarkerSymbol",_a,_5);_5.symbol.defaultPictureMarkerSymbol=_9;}return _a;});