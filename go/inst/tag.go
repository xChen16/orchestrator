/*
   Copyright 2015 Shlomi Noach, courtesy Booking.com

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package inst

import ()

type Tag struct {
}

type DatabaseInstanceTag struct {
	Key      *InstanceKey
	TagName  string
	TagValue string
	HasValue bool
}

func NewDatabaseInstanceTag(instanceKey *InstanceKey, tagName, tagValue string) *DatabaseInstanceTag {
	return &DatabaseInstanceTag{
		Key:      instanceKey,
		TagName:  tagName,
		TagValue: tagValue,
	}
}
