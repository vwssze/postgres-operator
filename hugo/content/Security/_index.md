---
title: "Security"
date:
draft: false
weight: 6
---

## Kubernetes RBAC

Install the requisite Operator RBAC resources, *as a Kubernetes cluster admin user*,  by running a Makefile target:

    make installrbac


This script creates the following RBAC resources on your Kubernetes cluster:

| Setting |Definition  |
|---|---|
| Custom Resource Definitions | pgbackups|
|  | pgclusters|
|  | pgpolicies|
|  | pgreplicas|
|  | pgtasks|
|  | pgupgrades|
| Cluster Roles | pgopclusterrole|
|  | pgopclusterrolecrd|
|  | scheduler-sa|
| Cluster Role Bindings | pgopclusterbinding|
|  | pgopclusterbindingcrd|
|  | scheduler-sa|
| Service Account | scheduler-sa|
| | postgres-operator|
| | pgo-backrest|
| | scheduler-sa|
| Roles| pgo-role|
| | pgo-backrest-role|
|Role Bindings | pgo-backrest-role-binding|





## Operator RBAC

The *conf/postgresql-operator/pgorole* file is read at start up time when the operator is deployed to the Kubernetes cluster.  This file defines the Operator roles whereby Operator API users can be authorized.

The *conf/postgresql-operator/pgouser* file is read at start up time also and contains username, password, and role information as follows:

    username:password:pgoadmin
    testuser:testpass:pgoadmin
    readonlyuser:testpass:pgoreader

A user creates a *.pgouser* file in their $HOME directory to identify
themselves to the Operator.  An entry in .pgouser will need to match
entries in the *conf/postgresql-operator/pgouser* file.  A sample
*.pgouser* file contains the following:

    username:password

The users pgouser file can also be located at:
*/etc/pgo/pgouser* or it can be found at a path specified by the
PGOUSER environment variable.

The following list shows the current complete list of possible pgo permissions:

|Permission|Description  |
|---|---|
|ApplyPolicy | allow *pgo apply*|
|CreateBackup | allow *pgo backup*|
|CreateCluster | allow *pgo create cluster*|
|CreateFailover | allow *pgo failover*|
|CreatePgbouncer | allow *pgo create pgbouncer*|
|CreatePgpool | allow *pgo create pgpool*|
|CreatePolicy | allow *pgo create policy*|
|CreateSchedule | allow *pgo create schedule*|
|CreateUpgrade | allow *pgo upgrade*|
|CreateUser | allow *pgo create user*|
|DeleteBackup | allow *pgo delete backup*|
|DeleteCluster | allow *pgo delete cluster*|
|DeletePgbouncer | allow *pgo delete pgbouncer*|
|DeletePgpool | allow *pgo delete pgpool*|
|DeletePolicy | allow *pgo delete policy*|
|DeleteSchedule | allow *pgo delete schedule*|
|DeleteUpgrade | allow *pgo delete upgrade*|
|DeleteUser | allow *pgo delete user*|
|DfCluster | allow *pgo df*|
|Label | allow *pgo label*|
|Load | allow *pgo load*|
|Reload | allow *pgo reload*|
|Restore | allow *pgo restore*|
|ShowBackup | allow *pgo show backup*|
|ShowCluster | allow *pgo show cluster*|
|ShowConfig | allow *pgo show config*|
|ShowPolicy | allow *pgo show policy*|
|ShowPVC | allow *pgo show pvc*|
|ShowSchedule | allow *pgo show schedule*|
|ShowUpgrade | allow *pgo show upgrade*|
|ShowWorkflow | allow *pgo show workflow*|
|Status | allow *pgo status*|
|TestCluster | allow *pgo test*|
|UpdateCluster | allow *pgo update cluster*|
|User | allow *pgo user*|
|Version | allow *pgo version*|


If the user is unauthorized for a pgo command, the user will
get back this response:

    FATA[0000] Authentication Failed: 40

## Making Security Changes
The Operator today requires you to make Operator security changes in the pgouser and pgorole files, and for those changes to take effect you are required to re-deploy the Operator:

    make deployoperator

This will recreate the *pgo-auth-secret* Secret that stores these files and is mounted by the Operator during its initialization.

## API Security
The Operator REST API is secured with keys stored in the *pgo-auth-secret* Secret.  Adjust the default keys to meet your security requirements using your own keys.  The *pgo-auth-secret* Secret is created when you run:

    make deployoperator

The keys are generated when the RBAC script is executed by the cluster admin:

    make installrbac
