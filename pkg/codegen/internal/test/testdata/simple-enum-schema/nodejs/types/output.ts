// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";

export interface Container {
    brightness?: enums.ContainerBrightness;
    color?: (enums.ContainerColor | string);
    material?: string;
    size: enums.ContainerSize;
}

