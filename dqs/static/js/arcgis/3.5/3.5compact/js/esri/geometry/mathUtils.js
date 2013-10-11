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
define("esri/geometry/mathUtils",["dojo/_base/lang","dojo/has","esri/kernel","esri/geometry/Point"],function(_1,_2,_3,_4){function _5(_6,_7){var dx=_7.x-_6.x,dy=_7.y-_6.y;return Math.sqrt(dx*dx+dy*dy);};function _8(_9,_a){var dx=_a[0]-_9[0],dy=_a[1]-_9[1];return Math.sqrt(dx*dx+dy*dy);};function _b(_c,_d,_e){if(_c instanceof _4){return new _4(_c.x+_e*(_d.x-_c.x),_c.y+_e*(_d.y-_c.y));}else{return [_c[0]+_e*(_d[0]-_c[0]),_c[1]+_e*(_d[1]-_c[1])];}};function _f(pt0,pt1){return _b(pt0,pt1,0.5);};function _10(n1,n2){return Math.abs(n1-n2)<1e-8;};function _11(p0,p1,p2,p3){var _12=10000000000,x,y,a0=_10(p0[0],p1[0])?_12:(p0[1]-p1[1])/(p0[0]-p1[0]),a1=_10(p2[0],p3[0])?_12:(p2[1]-p3[1])/(p2[0]-p3[0]),b0=p0[1]-a0*p0[0],b1=p2[1]-a1*p2[0];if(_10(a0,a1)){if(!_10(b0,b1)){return null;}else{if(_10(p0[0],p1[0])){if(Math.min(p0[1],p1[1])<Math.max(p2[1],p3[1])||Math.max(p0[1],p1[1])>Math.min(p2[1],p3[1])){y=(p0[1]+p1[1]+p2[1]+p3[1]-Math.min(p0[1],p1[1],p2[1],p3[1])-Math.max(p0[1],p1[1],p2[1],p3[1]))/2;x=(y-b0)/a0;}else{return null;}}else{if(Math.min(p0[0],p1[0])<Math.max(p2[0],p3[0])||Math.max(p0[0],p1[0])>Math.min(p2[0],p3[0])){x=(p0[0]+p1[0]+p2[0]+p3[0]-Math.min(p0[0],p1[0],p2[0],p3[0])-Math.max(p0[0],p1[0],p2[0],p3[0]))/2;y=a0*x+b0;}else{return null;}}return [x,y];}}if(_10(a0,_12)){x=p0[0];y=a1*x+b1;}else{if(_10(a1,_12)){x=p2[0];y=a0*x+b0;}else{x=-(b0-b1)/(a0-a1);y=a0*x+b0;}}return [x,y];};function _13(_14,_15,_16,_17,sr){var pt=_11([_14.x,_14.y],[_15.x,_15.y],[_16.x,_16.y],[_17.x,_17.y]);if(pt){pt=new _4(pt[0],pt[1],sr);}return pt;};function _18(_19,_1a){var p1=_19[0],p2=_19[1],p3=_1a[0],p4=_1a[1],x1=p1[0],y1=p1[1],x2=p2[0],y2=p2[1],x3=p3[0],y3=p3[1],x4=p4[0],y4=p4[1],x43=x4-x3,x13=x1-x3,x21=x2-x1,y43=y4-y3,y13=y1-y3,y21=y2-y1,_1b=(y43*x21)-(x43*y21),ua,ub,px,py;if(_1b===0){return false;}ua=((x43*y13)-(y43*x13))/_1b;ub=((x21*y13)-(y21*x13))/_1b;if(ua>=0&&ua<=1&&ub>=0&&ub<=1){px=x1+(ua*(x2-x1));py=y1+(ua*(y2-y1));return [px,py];}else{return false;}};function _1c(_1d,_1e){var p1=_1e[0],p2=_1e[1],x1=p1[0],y1=p1[1],x2=p2[0],y2=p2[1],x3=_1d[0],y3=_1d[1],x21=x2-x1,y21=y2-y1,x31=x3-x1,y31=y3-y1,_1f=Math.sqrt,pow=Math.pow,mag=_1f(pow(x21,2)+pow(y21,2)),u=((x31*x21)+(y31*y21))/(mag*mag),x=x1+u*x21,y=y1+u*y21;return _1f(pow(x3-x,2)+pow(y3-y,2));};var _20={getLength:_5,_getLength:_8,getPointOnLine:_b,getMidpoint:_f,_equals:_10,_getLineIntersection:_11,getLineIntersection:_13,_getLineIntersection2:_18,_pointLineDistance:_1c};if(_2("extend-esri")){_1.mixin(_1.getObject("geometry",true,_3),_20);}return _20;});