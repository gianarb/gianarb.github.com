---
layout: post
title:  "Staging environment on demand with AWS Cloudformation"
date:   2015-07-08 09:08:27
img: /img/amazon-aws-logo.jpg
categories: [post]
tags: [devops, aws]
summary: Build your stagin environment on demand with AWS Cloudformation
priority: 0.6
changefreq: yearly
---

<blockquote class="twitter-tweet tw-align-center" lang="en"><p lang="en" dir="ltr">Staging environment on demand. To Work on <a href="https://twitter.com/hashtag/AWS?src=hash">#AWS</a> low level with <a href="https://twitter.com/hashtag/cloudformation?src=hash">#cloudformation</a> <a href="http://t.co/VWBR129637">http://t.co/VWBR129637</a> <a href="https://twitter.com/hashtag/cloud?src=hash">#cloud</a> <a href="https://twitter.com/hashtag/devops?src=hash">#devops</a></p>&mdash; Gianluca Arbezzano (@GianArb) <a href="https://twitter.com/GianArb/status/621691855810494464">July 16, 2015</a></blockquote>
<script async src="//platform.twitter.com/widgets.js" charset="utf-8"></script>

## Staging Environment
There are few environments during my developer workflow, today I chose a little example:

* Production enviroment always exists, it runs the stable application and you can not use it for your test.
* **Staging** enviroment is a "pre-production" state.
* Develop enviroment is instable and it runs new features and fixes, here there's the work of all team but it's not ready to go in production.

![Staging graph](/img/cloudformation-staging/staging.jpg)

Staging environment in my opinion could be "volatile" version, we use it when our product is ready to go in production for the last time it was unused. Maybe this statement isn't real in your work but if you think a little team of consultants that
 work on different projects maybe this words have a sense.

## AWS Cloudformation
CloudFormation is an AWS service that helps you to orchestate all AWS services, you can write a template in JSON and you can use it to create an infrastructure with one click.
This solution helps me  to build and destroy this environment and we can pay it only if it's necessary, if you use a `stagin env == production env` it can be very very expensive.
This solution could help you to down cost.


## Current infrastructure

![RDS and EC2 infrastructure](/img/cloudformation-staging/infra.jpg)

This is my template to build a simple application Frontend + MySQL (RDS).
In this implementation I build network configuration and I create one instance of RDS and one EC2 (my frontend).
`Parameters` key is the list of external parameters that I can use to configure my template, for example database and EC2 key pair, my root's password..
`Resources` key contains description of all actors of this infrastructure.

{% highlight json %}
{
  "Parameters" : {
    "VPCName" : {
      "Type" : "String",
      "Default" : "staging",
      "Description" : "VPC name"
    },
    "ProjectName" : {
      "Type" : "String",
      "Default" : "app",
      "Description" : "Project name"
    },
    "WebKey" : {
      "Type" : "String",
      "Default" : "web-key",
      "Description" : "Ssh key to log into the web instances"
    },
    "WebInstanceType" : {
      "Type" : "String",
      "Default" : "m3.medium",
      "Description" : "Web instance type"
    },
    "WebInstanceImage" : {
      "Type" : "String",
      "Default" : "ami-47a23a30",
      "Description" : "Web instance image"
    },
    "DatabaseInstanceType" : {
      "Type" : "String",
      "Default" : "db.m3.medium",
      "Description" : "Database instance type"
    },
    "DatabaseName" : {
      "Type" : "String",
      "Default" : "mydb",
      "Description" : "Database instance's name"
    },
    "DatabaseMasterUsername" : {
      "Type" : "String",
      "Default" : "gianarb",
      "Description" : "Name of master user"
    },
    "DatabaseEngineVersion" : {
      "Type" : "String",
      "Default" : "5.6",
      "Description" : "MySQL version"
    },
    "DatabaseUserPassword" : {
      "Type" : "String",
      "Default" : "test1234",
      "Description" : "User password"
    },
    "DatabasePublicAccess" : {
      "Type" : "String",
      "Default" : true
    },
    "DatabaseMultiAZ" : {
      "Type" : "String",
      "Default" : false
    }
  },
  "Resources" : {
    "Staging": {
       "Type" : "AWS::EC2::VPC",
       "Properties" : {
          "CidrBlock" : "10.15.0.0/16",
          "EnableDnsSupport" : true,
          "EnableDnsHostnames" : true,
          "InstanceTenancy" : "default",
          "Tags" : [{"Key": "Name", "Value": {"Ref": "VPCName"}}]
       }
    },
    "DatabaseSubnet1": {
      "Type" : "AWS::EC2::Subnet",
      "Properties" : {
        "AvailabilityZone" : "eu-west-1a",
        "CidrBlock" : "10.15.1.0/28",
        "MapPublicIpOnLaunch" : true,
        "VpcId": {
          "Ref" : "Staging"
        },
        "Tags": [{"Key": "Name", "Value": "db-1a"}]
      }
    },
    "DatabaseSubnet2": {
      "Type" : "AWS::EC2::Subnet",
      "Properties" : {
        "AvailabilityZone" : "eu-west-1b",
        "CidrBlock" : "10.15.1.16/28",
        "MapPublicIpOnLaunch" : true,
        "VpcId": {
          "Ref" : "Staging"
        },
        "Tags" : [{"Key": "Name", "Value": "db-1b"}]
      }
    },
    "WebSubnet1": {
      "Type" : "AWS::EC2::Subnet",
      "Properties" : {
        "AvailabilityZone" : "eu-west-1a",
        "CidrBlock" : "10.15.0.8/28",
        "MapPublicIpOnLaunch" : true,
        "VpcId": {
          "Ref" : "Staging"
        },
        "Tags" : [{"Key": "Name", "Value": "web-1a"}]
      }
    },
    "RDSSubnet": {
     "Type" : "AWS::RDS::DBSubnetGroup",
     "Properties" : {
        "DBSubnetGroupDescription": "db-prod-subnet-group",
        "SubnetIds" : [
          { "Ref": "DatabaseSubnet1" },
          { "Ref": "DatabaseSubnet2" }
        ]
      }
    },
    "Database": {
      "Type" : "AWS::RDS::DBInstance",
      "Properties" : {
        "AllocatedStorage": "5",
        "AllowMajorVersionUpgrade" : false,
        "DBInstanceClass": {"Ref":"DatabaseInstanceType"},
        "DBName" : {"Ref":"DatabaseName"},
        "DBInstanceIdentifier": {"Ref":"DatabaseName"},
        "Engine" : "MySQL",
        "EngineVersion" : {"Ref":"DatabaseEngineVersion"},
        "DBSubnetGroupName": {
          "Ref": "RDSSubnet"
        },
        "MasterUsername" : {"Ref": "DatabaseMasterUsername"},
        "MasterUserPassword" : {"Ref": "DatabaseUserPassword"},
        "MultiAZ" : true,
        "VPCSecurityGroups": [
          {
            "Ref": "DatabaseSG"
          }
        ],
        "PubliclyAccessible" : {"Ref": "DatabasePublicAccess"},
        "Tags" : [{"Key": "Name", "Value": {"Fn::Join":[".", ["db", {"Ref": "ProjectName"}, {"Ref":"VPCName"}]]} }]
      }
    },
    "WebInstance" : {
        "Type" : "AWS::EC2::Instance",
        "Properties" : {
            "ImageId" : {"Ref": "WebInstanceImage"},
            "InstanceType" : {"Ref": "WebInstanceType"},
            "KeyName" : {"Ref": "WebKey"},
            "BlockDeviceMappings" : [
                {
                    "DeviceName" : "/dev/sdm",
                    "Ebs" : {
                        "VolumeType" : "io1",
                        "Iops" : "200",
                        "DeleteOnTermination" : "false",
                        "VolumeSize" : "20"
                    }
                },
                {
                    "DeviceName" : "/dev/sdk",
                    "NoDevice" : {}
                 }
            ],
            "SubnetId": { "Ref" : "WebSubnet1" },
            "SecurityGroupIds": [
                {"Ref": "WebSG"}
            ]
        }
    },
    "StagingZone": {
      "Type" : "AWS::Route53::HostedZone",
      "Properties" : {
        "Name" : {"Fn::Join":[".", [{"Ref": "ProjectName"}, {"Ref":"VPCName"}]]},
        "VPCs" : [{"VPCId": {"Ref": "Staging"}, "VPCRegion": "eu-west-1"}]
      }
    },
    "StagingInternetGateway" : {
      "Type" : "AWS::EC2::InternetGateway",
      "Properties" : {
        "Tags" : [ {"Key" : "Name", "Value" : {"Fn::Join":["-", [{"Ref":"VPCName"}, "igw"]]}}]
      }
    },
    "StagingIgwAttach": {
      "Type" : "AWS::EC2::VPCGatewayAttachment",
      "Properties" : {
        "InternetGatewayId" : {"Ref": "StagingInternetGateway"},
        "VpcId" : {"Ref": "Staging"}
      }
    },
    "StagingRouteTable": {
       "Type" : "AWS::EC2::RouteTable",
       "Properties" : {
          "VpcId" : {"Ref": "Staging"}
       }
    },
    "LocalRoute": {
       "Type" : "AWS::EC2::Route",
       "Properties" : {
          "DestinationCidrBlock" : "0.0.0.0/0",
          "GatewayId" : {"Ref": "StagingInternetGateway"},
          "RouteTableId" : {"Ref": "StagingRouteTable"}
       }
    },
    "Web1LocalRoute": {
      "Type" : "AWS::EC2::SubnetRouteTableAssociation",
      "Properties" : {
        "RouteTableId" : {"Ref": "StagingRouteTable"},
        "SubnetId" : {"Ref": "WebSubnet1"}
      }
    },
    "Db1LocalRoute": {
      "Type" : "AWS::EC2::SubnetRouteTableAssociation",
      "Properties" : {
        "RouteTableId" : {"Ref": "StagingRouteTable"},
        "SubnetId" : {"Ref": "DatabaseSubnet1"}
      }
    },
    "Db2LocalRoute": {
      "Type" : "AWS::EC2::SubnetRouteTableAssociation",
      "Properties" : {
        "RouteTableId" : {"Ref": "StagingRouteTable"},
        "SubnetId" : {"Ref": "DatabaseSubnet2"}
      }
    },
    "DatabaseSG": {
      "Type" : "AWS::EC2::SecurityGroup",
      "Properties" : {
        "GroupDescription" : "Database security groups",
        "SecurityGroupIngress" : [
          {
            "IpProtocol" : "tcp",
            "FromPort": 3306,
            "ToPort" : "3306",
            "SourceSecurityGroupId": {"Ref" : "WebSG"}
          }
        ],
        "Tags" :  [{"Key": "Name", "Value": "db-sg"}],
        "VpcId" : {"Ref": "Staging"}
      }
    },
    "WebSG": {
      "Type" : "AWS::EC2::SecurityGroup",
      "Properties" : {
        "GroupDescription" : "Web security groups",
        "SecurityGroupIngress" : [
          {
            "IpProtocol" : "tcp",
            "ToPort" : 80,
            "FromPort": 80,
            "CidrIp" : "0.0.0.0/0"
          },
          {
            "IpProtocol" : "tcp",
            "ToPort" : 22,
            "FromPort": 22,
            "CidrIp" : "0.0.0.0/0"
          }
        ],
        "Tags" :  [{"Key": "Name", "Value": "web-sg"}],
        "VpcId" : {"Ref": "Staging"}
      }
    },
    "DatabaseRecordSet" : {
      "Type" : "AWS::Route53::RecordSet",
      "Properties" : {
         "HostedZoneId" : {
            "Ref": "StagingZone"
         },
         "Comment" : "DNS name for database",
         "Name" : {"Fn::Join":[".", ["db", {"Ref": "ProjectName"}, {"Ref":"VPCName"}]]},
         "Type" : "CNAME",
         "TTL" : "300",
         "ResourceRecords" : [
           { "Fn::GetAtt" : [ "Database", "Endpoint.Address"]}
         ]
      }
    }
  }
}
{% endhighlight %}

## Conclusion
You can load this teamplate in your account and after environment creations you are ready to work with one EC2 instance and one RDS with MySQL 5.6 installed.
You can log into the web interface with key-pair chosen during the creation flow (default ga-eu) and I set default this mysql credential:

* user gianarb
* password test1234

But you can change it before running this template because they are `Parameters`.
This approach in my opinion is very powerful because you can start versioning your infrastructure and you can delete and restore it quickly because if you delete the cloudformation stack it rollbacks all resources, it is very easy!

## Trick

Parameters node create a form into the AWS CloudFormation console to choose a lot of different variable values, for example name of intances or key-pair to log in your EC2.

{% highlight json %}
{
  "Parameters" : {
    "VPCName" : {
      "Type" : "String",
      "Default" : "staging",
      "Description" : "VPC name"
    },
    "ProjectName" : {
      "Type" : "String",
      "Default" : "app",
      "Description" : "Project name"
    },
    "WebKey" : {
      "Type" : "String",
      "Default" : "web-key",
      "Description" : "Ssh key to log into the web instances"
    }
}
{% endhighlight %}

<hr class="style-two">

Resources node contains all elements of your infrastructure, EC2, RDS, VCP.. You can use the parameteters with a simple `Ref Key`.
es. `[{"Key": "Name", "Value": "ProjectName"}]` describe the name of the specific project into the parameter form.

{% highlight json %}
{
  "Resources" : {
    "Staging": {
       "Type" : "AWS::EC2::VPC",
       "Properties" : {
          "CidrBlock" : "10.15.0.0/16",
          "EnableDnsSupport" : true,
          "EnableDnsHostnames" : true,
          "InstanceTenancy" : "default",
          "Tags" : [{"Key": "Name", "Value": {"Ref": "VPCName"}}]
       }
    },
    "DatabaseSubnet1": {
      "Type" : "AWS::EC2::Subnet",
      "Properties" : {
        "AvailabilityZone" : "eu-west-1a",
        "CidrBlock" : "10.15.1.0/28",
        "MapPublicIpOnLaunch" : true,
        "VpcId": {
          "Ref" : "Staging"
        },
        "Tags": [{"Key": "Name", "Value": "db-1a"}]
      }
    }
}
{% endhighlight %}

<hr class="style-two">

In your template you can describe VPC and create its subnet. You can also describe the specific resource and you can use it to build another
{% highlight json %}
WebSubnet1": {
  "Type" : "AWS::EC2::Subnet",
  "Properties" : {
    "AvailabilityZone" : "eu-west-1a",
    "CidrBlock" : "10.15.0.8/28",
    "MapPublicIpOnLaunch" : true,
    "VpcId": {
      "Ref" : "Staging"
    },
    "Tags" : [{"Key": "Name", "Value": "web-1a"}]
  }
},
{% endhighlight %}
In this example I resumed `Staging` VPC to build its subnet.

<hr class="style-two">

This chapter is insteresting because it creates a RecordSet to map a CNAME DNS in your VPC and now in your Web instances you can resolve MYSql host with `db.app.staging`.

{% highlight json %}
"DatabaseRecordSet" : {
  "Type" : "AWS::Route53::RecordSet",
  "Properties" : {
     "HostedZoneId" : {
        "Ref": "StagingZone"
     },
     "Comment" : "DNS name for database",
     "Name" : {"Fn::Join":[".", ["db", {"Ref": "ProjectName"}, {"Ref":"VPCName"}]]},
     "Type" : "CNAME",
     "TTL" : "300",
     "ResourceRecords" : [
       { "Fn::GetAtt" : [ "Database", "Endpoint.Address"]}
     ]
  }
}
{% endhighlight %}

<br/>
<br/>
<br/>

<div class="well"><a target="_blank" href="https://twitter.com/EmanueleMinotto">@EmanualeMinotto</a> thanks for trying to fix my bad English</div>
