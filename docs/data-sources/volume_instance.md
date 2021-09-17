---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "zedcloud_volume_instance Data Source - terraform-provider-zedcloud"
subcategory: ""
description: |-
  Schema for data source zedcloudvolumeinstance. Must specify id or name
---

# zedcloud_volume_instance (Data Source)

Schema for data source zedcloud_volume_instance. Must specify id or name



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **name** (String) User defined name of the object. Must be unique across the enterprise. Once object is created, name can’t be changed

### Optional

- **accessmode** (String) Access mode. Valid Values: VOLUME_INSTANCE_ACCESS_MODE_INVALID, VOLUME_INSTANCE_ACCESS_MODE_READWRITE, VOLUME_INSTANCE_ACCESS_MODE_READONLY, VOLUME_INSTANCE_ACCESS_MODE_MULTIREAD_SINGLEWRITE
- **cleartext** (Boolean) Flag to keep the contents of the volume unencrypted (in clear text)
- **content_tree_id** (String) Content tree ID
- **description** (String) Detailed description of the Object
- **device_id** (String) ID of the device on which volume instance is created
- **image** (String) Name of the image
- **implicit** (Boolean) Flag to create implicit volumes
- **label** (String) Label
- **multiattach** (Boolean) Flag to enable the volume to be attached to multiple app instances
- **project_id** (String) ID of the project to which the Object belongs
- **revision** (Block List, Max: 1) System defined revision information of the object (see [below for nested schema](#nestedblock--revision))
- **size_bytes** (String) Size of volume
- **title** (String) User defined title of the object. title can be changed any time.
- **type** (String) Type of Volume Instance. Valid Values: VOLUME_INSTANCE_TYPE_UNSPECIFIED, VOLUME_INSTANCE_TYPE_EMPTYDIR, VOLUME_INSTANCE_TYPE_BLOCKSTORAGE, VOLUME_INSTANCE_TYPE_HOSTFS, VOLUME_INSTANCE_TYPE_TMPFS, VOLUME_INSTANCE_TYPE_SECRET, VOLUME_INSTANCE_TYPE_NFS, VOLUME_INSTANCE_TYPE_AWS_BLOCK_STORAGE, VOLUME_INSTANCE_TYPE_CONTENT_TREE

### Read-Only

- **id** (String) System defined unique Id of the Object
- **purge** (List of Object) Purge Counter information (see [below for nested schema](#nestedatt--purge))

<a id="nestedblock--revision"></a>
### Nested Schema for `revision`

Optional:

- **created_at** (String) The time, in milliseconds since the epoch, when the object was created
- **created_by** (String) User who Created the object
- **curr** (String) Current version of the object
- **prev** (String) Prev version of the object
- **updated_at** (String) The time, in milliseconds since the epoch, when the object was last updated
- **updated_by** (String) User who last updated the object


<a id="nestedatt--purge"></a>
### Nested Schema for `purge`

Read-Only:

- **counter** (Number)
- **ops_time** (String)

