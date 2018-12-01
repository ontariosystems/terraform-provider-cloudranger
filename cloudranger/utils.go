package cloudranger

import (
	"github.com/hashicorp/terraform/helper/schema"
	cr "github.com/jaxi/terraform-provider-cloudranger/client"
)

func tagsFrom(d *schema.ResourceData) []cr.Tag {
	tags := []cr.Tag{}
	if attr, ok := d.GetOk("tags"); ok {
		for _, tagItem := range attr.([]interface{}) {
			var tag cr.Tag
			tagDataSrc := tagItem.(map[string]interface{})
			tagKey, ok := tagDataSrc["key"]
			if !ok {
				continue
			}
			tag.Key = tagKey.(string)

			tagValue, ok := tagDataSrc["value"]
			if !ok {
				continue
			}
			tag.Value = tagValue.(string)

			tags = append(tags, tag)
		}
	}

	return tags
}

func loadTags(tags []cr.Tag, d *schema.ResourceData) error {
	tagsData := [](map[string]string){}

	for _, tag := range tags {
		tagsData = append(tagsData, map[string]string{
			"key":   tag.Key,
			"value": tag.Value,
		})
	}
	if err := d.Set("tags", tagsData); err != nil {
		return err
	}

	return nil
}
