/*
Package authentication is the client.Device.AuthenticationProfile namespace.

For Panorama, there are two possibilities:  managing this object on Panorama
itself or inside of a Template.

To manage objects save on Panorama, leave "tmpl", "ts", and "vsys" params empty.

To manage objects in a template, specify the template name and the vsys (if
unspecified, defaults to "shared").

Normalized object:  Entry
*/
package authentication
