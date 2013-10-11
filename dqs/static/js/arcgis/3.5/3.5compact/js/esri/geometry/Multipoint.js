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
define("esri/geometry/Multipoint",["dojo/_base/declare","dojo/_base/lang","dojo/has","esri/kernel","esri/SpatialReference","esri/geometry/Geometry","esri/geometry/Point","esri/geometry/Extent"],function(_1,_2,_3,_4,_5,_6,_7,_8){var _9={type:"multipoint",points:null};var _a=_1(_6,{declaredClass:"esri.geometry.Multipoint",constructor:function(_b){_2.mixin(this,_9);this.points=[];if(_b){if(_b.points){_2.mixin(this,_b);}else{this.spatialReference=_b;}if(this.spatialReference){this.spatialReference=new _5(this.spatialReference);}}this.verifySR();},_extent:null,addPoint:function(_c){this._extent=null;if(_2.isArray(_c)){this.points.push(_c);}else{this.points.push([_c.x,_c.y]);}return this;},removePoint:function(_d){if(this._validateInputs(_d)){this._extent=null;return new _7(this.points.splice(_d,1)[0],this.spatialReference);}},getExtent:function(){if(this._extent){return new _8(this._extent);}var _e=this.points,il=_e.length;if(!il){return;}var _f=_e[0],_10,_11,_12=(_10=_f[0]),_13=(_11=_f[1]),min=Math.min,max=Math.max,sr=this.spatialReference,x,y,i;for(i=0;i<il;i++){_f=_e[i];x=_f[0];y=_f[1];_12=min(_12,x);_13=min(_13,y);_10=max(_10,x);_11=max(_11,y);}this._extent={xmin:_12,ymin:_13,xmax:_10,ymax:_11,spatialReference:sr?sr.toJson():null};return new _8(this._extent);},_validateInputs:function(_14){if(_14===null||_14<0||_14>=this.points.length){return false;}return true;},getPoint:function(_15){if(this._validateInputs(_15)){var _16=this.points[_15];return new _7(_16[0],_16[1],this.spatialReference);}},setPoint:function(_17,_18){if(this._validateInputs(_17)){this._extent=null;this.points[_17]=[_18.x,_18.y];return this;}},toJson:function(){var _19={points:_2.clone(this.points)},sr=this.spatialReference;if(sr){_19.spatialReference=sr.toJson();}return _19;}});_a.defaultProps=_9;if(_3("extend-esri")){_2.setObject("geometry.Multipoint",_a,_4);_4.geometry.defaultMultipoint=_9;}return _a;});