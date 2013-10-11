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
define("esri/virtualearth/VEGeocodeResult",["dojo/_base/declare","dojo/_base/lang","dojo/has","esri/kernel","esri/geometry/Point","esri/geometry/Extent","esri/virtualearth/VEAddress"],function(_1,_2,_3,_4,_5,_6,_7){var _8=_1(null,{declaredClass:"esri.virtualearth.VEGeocodeResult",constructor:function(_9){_2.mixin(this,{address:null,bestView:null,calculationMethod:null,confidence:null,displayName:null,entityType:null,location:null,matchCodes:null},_9);if(this.address){this.address=new _7(this.address);}if(this.bestView){this.bestView=new _6(this.bestView);}if(this.locationArray){this.calculationMethod=this.locationArray[0].calculationMethod;this.location=new _5(this.locationArray[0]);}}});if(_3("extend-esri")){_2.setObject("virtualearth.VEGeocodeResult",_8,_4);}return _8;});