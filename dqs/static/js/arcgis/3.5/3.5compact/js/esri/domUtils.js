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
define("esri/domUtils",["dojo/_base/connect","dojo/_base/lang","dojo/dom-style","dojo/has","esri/kernel"],function(_1,_2,_3,_4,_5){var _6={show:function(_7){if(_7){_7.style.display="block";}},hide:function(_8){if(_8){_8.style.display="none";}},toggle:function(_9){_9.style.display=_9.style.display==="none"?"block":"none";},documentBox:_4("ie")?{w:document.documentElement.clientWidth,h:document.documentElement.clientHeight}:{w:window.innerWidth,h:window.innerHeight},setScrollable:function(_a){var _b=0,_c=0,_d=0,_e=0,_f=0,_10=0;return [_1.connect(_a,"ontouchstart",function(evt){_b=evt.touches[0].screenX;_c=evt.touches[0].screenY;_d=_a.scrollWidth;_e=_a.scrollHeight;_f=_a.clientWidth;_10=_a.clientHeight;}),_1.connect(_a,"ontouchmove",function(evt){evt.preventDefault();var _11=_a.firstChild;if(_11 instanceof Text){_11=_a.childNodes[1];}var _12=_11._currentX||0,_13=_11._currentY||0;_12+=(evt.touches[0].screenX-_b);if(_12>0){_12=0;}else{if(_12<0&&(Math.abs(_12)+_f)>_d){_12=-1*(_d-_f);}}_11._currentX=_12;_13+=(evt.touches[0].screenY-_c);if(_13>0){_13=0;}else{if(_13<0&&(Math.abs(_13)+_10)>_e){_13=-1*(_e-_10);}}_11._currentY=_13;_3.set(_11,{"-webkit-transition-property":"-webkit-transform","-webkit-transform":"translate("+_12+"px, "+_13+"px)"});_b=evt.touches[0].screenX;_c=evt.touches[0].screenY;})];}};if(_4("extend-esri")){_2.mixin(_5,_6);}return _6;});