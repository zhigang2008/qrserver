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
define("esri/tasks/JobInfo",["dojo/_base/declare","dojo/_base/lang","dojo/has","esri/kernel","esri/tasks/GPMessage"],function(_1,_2,_3,_4,_5){var _6=_1(null,{declaredClass:"esri.tasks.JobInfo",constructor:function(_7){this.messages=[];_2.mixin(this,_7);var _8=this.messages,i,il=_8.length;for(i=0;i<il;i++){_8[i]=new _5(_8[i]);}},jobId:"",jobStatus:""});_2.mixin(_6,{STATUS_CANCELLED:"esriJobCancelled",STATUS_CANCELLING:"esriJobCancelling",STATUS_DELETED:"esriJobDeleted",STATUS_DELETING:"esriJobDeleting",STATUS_EXECUTING:"esriJobExecuting",STATUS_FAILED:"esriJobFailed",STATUS_NEW:"esriJobNew",STATUS_SUBMITTED:"esriJobSubmitted",STATUS_SUCCEEDED:"esriJobSucceeded",STATUS_TIMED_OUT:"esriJobTimedOut",STATUS_WAITING:"esriJobWaiting"});if(_3("extend-esri")){_2.setObject("tasks.JobInfo",_6,_4);}return _6;});