// *** WARNING: this file was generated by the Coconut IDL Compiler (CIDLC).  ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as coconut from "@coconut/coconut";

import {Function} from "./function";

export class Watch extends coconut.Resource implements WatchArgs {
    public readonly name: string;
    public namespace: string;
    public objType: string;
    public labelSelector: string;
    public fieldSelector: string;
    public function: Function;
    public target: string;

    constructor(args: WatchArgs) {
        super();
        if (args.name === undefined) {
            throw new Error("Missing required argument 'name'");
        }
        this.name = args.name;
        if (args.namespace === undefined) {
            throw new Error("Missing required argument 'namespace'");
        }
        this.namespace = args.namespace;
        if (args.objType === undefined) {
            throw new Error("Missing required argument 'objType'");
        }
        this.objType = args.objType;
        if (args.labelSelector === undefined) {
            throw new Error("Missing required argument 'labelSelector'");
        }
        this.labelSelector = args.labelSelector;
        if (args.fieldSelector === undefined) {
            throw new Error("Missing required argument 'fieldSelector'");
        }
        this.fieldSelector = args.fieldSelector;
        if (args.function === undefined) {
            throw new Error("Missing required argument 'function'");
        }
        this.function = args.function;
        if (args.target === undefined) {
            throw new Error("Missing required argument 'target'");
        }
        this.target = args.target;
    }
}

export interface WatchArgs {
    readonly name: string;
    namespace: string;
    objType: string;
    labelSelector: string;
    fieldSelector: string;
    function: Function;
    target: string;
}


