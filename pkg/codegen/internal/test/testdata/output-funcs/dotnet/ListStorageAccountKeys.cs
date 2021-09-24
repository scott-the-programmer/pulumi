// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.Mypkg
{
    public static class ListStorageAccountKeys
    {
        /// <summary>
        /// The response from the ListKeys operation.
        /// API Version: 2021-02-01.
        /// </summary>
        public static Task<ListStorageAccountKeysResult> InvokeAsync(ListStorageAccountKeysArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<ListStorageAccountKeysResult>("mypkg::listStorageAccountKeys", args ?? new ListStorageAccountKeysArgs(), options.WithVersion());

        /// <summary>
        /// The response from the ListKeys operation.
        /// API Version: 2021-02-01.
        /// </summary>
        public static Output<ListStorageAccountKeysResult> Invoke(ListStorageAccountKeysInvokeArgs args, InvokeOptions? options = null)
        {
            return Pulumi.Output.All(
                args.AccountName.Box(),
                args.Expand.Box(),
                args.ResourceGroupName.Box()
            ).Apply(a =>
            {
                var args = new ListStorageAccountKeysArgs();
                a[0].Set(args, nameof(args.AccountName));
                a[1].Set(args, nameof(args.Expand));
                a[2].Set(args, nameof(args.ResourceGroupName));
                return InvokeAsync(args, options);
            });
        }
    }


    public sealed class ListStorageAccountKeysArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
        /// </summary>
        [Input("accountName", required: true)]
        public string AccountName { get; set; } = null!;

        /// <summary>
        /// Specifies type of the key to be listed. Possible value is kerb.
        /// </summary>
        [Input("expand")]
        public string? Expand { get; set; }

        /// <summary>
        /// The name of the resource group within the user's subscription. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public ListStorageAccountKeysArgs()
        {
        }
    }

    public sealed class ListStorageAccountKeysInvokeArgs
    {
        /// <summary>
        /// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
        /// </summary>
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// Specifies type of the key to be listed. Possible value is kerb.
        /// </summary>
        public Input<string>? Expand { get; set; }

        /// <summary>
        /// The name of the resource group within the user's subscription. The name is case insensitive.
        /// </summary>
        public Input<string> ResourceGroupName { get; set; } = null!;

        public ListStorageAccountKeysInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class ListStorageAccountKeysResult
    {
        /// <summary>
        /// Gets the list of storage account keys and their properties for the specified storage account.
        /// </summary>
        public readonly ImmutableArray<Outputs.StorageAccountKeyResponse> Keys;

        [OutputConstructor]
        private ListStorageAccountKeysResult(ImmutableArray<Outputs.StorageAccountKeyResponse> keys)
        {
            Keys = keys;
        }
    }
}
