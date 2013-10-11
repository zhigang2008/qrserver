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
define("esri/graphic",["dojo/_base/declare","dojo/_base/lang","dojo/has","esri/kernel","esri/domUtils","esri/lang","esri/InfoTemplate","esri/geometry/jsonUtils","esri/symbols/jsonUtils"],function(_1,_2,_3,_4,_5,_6,_7,_8,_9){var _a=_1(null,{declaredClass:"esri.Graphic",constructor:function(_b,_c,_d,_e){if(_b&&!_b.declaredClass){this.geometry=_b.geometry?_8.fromJson(_b.geometry):null;this.symbol=_b.symbol?_9.fromJson(_b.symbol):null;this.attributes=_b.attributes||null;this.infoTemplate=_b.infoTemplate?new _7(_b.infoTemplate):null;}else{this.geometry=_b;this.symbol=_c;this.attributes=_d;this.infoTemplate=_e;}},_shape:null,_graphicsLayer:null,_visible:true,visible:true,getDojoShape:function(){return this._shape;},getLayer:function(){return this._graphicsLayer;},setGeometry:function(_f){this.geometry=_f;var gl=this._graphicsLayer;if(gl){gl._updateExtent(this);gl._draw(this,true);}return this;},setSymbol:function(_10,_11){var gl=this._graphicsLayer,_12=this._shape;this.symbol=_10;if(_10){this.symbol._stroke=this.symbol._fill=null;}if(gl){if(_11){if(_12){gl._removeShape(this);}}gl._draw(this,true);}return this;},setAttributes:function(_13){this.attributes=_13;return this;},setInfoTemplate:function(_14){this.infoTemplate=_14;return this;},getInfoTemplate:function(){return this._getEffInfoTemplate();},_getEffInfoTemplate:function(){var _15=this.getLayer();return this.infoTemplate||(_15&&_15.infoTemplate);},getTitle:function(){var _16=this.getInfoTemplate();var _17=_16&&_16.title;if(_2.isFunction(_17)){_17=_17.call(_16,this);}else{if(_2.isString(_17)){var _18=this._graphicsLayer;var _19=_18&&_18._getDateOpts;_17=_6.substitute(this.attributes,_17,{first:true,dateFormat:_19&&_19.call(_18)});}}return _17;},getContent:function(){var _1a=this.getInfoTemplate();var _1b=_1a&&_1a.content;if(_2.isFunction(_1b)){_1b=_1b.call(_1a,this);}else{if(_2.isString(_1b)){var _1c=this._graphicsLayer;var _1d=_1c&&_1c._getDateOpts;_1b=_6.substitute(this.attributes,_1b,{dateFormat:_1d&&_1d.call(_1c)});}}return _1b;},show:function(){this.visible=this._visible=true;var _1e=this._shape,_1f;if(_1e){_1f=_1e.declaredClass.toLowerCase().indexOf("canvas")===-1?_1e.getEventSource():null;if(_1f){_5.show(_1f);}}else{if(this._graphicsLayer){this._graphicsLayer._draw(this,true);}}return this;},hide:function(){this.visible=this._visible=false;var _20=this._shape,_21,_22;if(_20){_21=_20.declaredClass.toLowerCase().indexOf("canvas")===-1?_20.getEventSource():null;if(_21){_5.hide(_21);}else{_22=this._graphicsLayer;if(_22){_22._removeShape(this);}}}return this;},toJson:function(){var _23={};if(this.geometry){_23.geometry=this.geometry.toJson();}if(this.attributes){_23.attributes=_2.mixin({},this.attributes);}if(this.symbol){_23.symbol=this.symbol.toJson();}if(this.infoTemplate){_23.infoTemplate=this.infoTemplate.toJson();}return _23;}});if(_3("extend-esri")){_4.Graphic=_a;}return _a;});