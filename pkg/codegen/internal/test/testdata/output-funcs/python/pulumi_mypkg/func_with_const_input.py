# coding=utf-8
# *** WARNING: this file was generated by test. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'func_with_const_input',
]

def func_with_const_input(plain_input: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None):
    """
    Codegen demo with const inputs
    """
    __args__ = dict()
    __args__['plainInput'] = plain_input
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('mypkg::funcWithConstInput', __args__, opts=opts).value

