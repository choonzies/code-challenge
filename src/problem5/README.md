 ### What is consensus-breaking change
 * A consensus-breaking change refers to a modification in a system that disrupts the agreement or consistency among participants or nodes. 
 * Such changes can lead to incompatibility issues, requiring all participants to update their systems to maintain consensus.
 
 ### My implementation
 * In this case, I adjusted the field type in `post.proto` (template for "resource") from `string` to `uint64`. This change breaks consensus because it alters the data type expected by the system, potentially causing issues with data interpretation and processing across different nodes or participants.
