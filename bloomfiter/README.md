## Bloom Filter
A Bloom filter is a space-efficient data structure used to test whether an element is a member of a set. 
It can have false positives but no false negatives.

**Pros**: Can store additional information beyond membership, efficient space utilization.
**Cons**: Higher false positive rate, limited support for deletion.

### Use Cases:

**Web caching**: Web caching systems like Squid use Bloom filters to check if a requested URL is already cached, avoiding unnecessary disk lookups.
Database systems: Bloom filters can be used to speed up query processing in databases by quickly determining if a specific record exists in a large dataset without scanning the entire dataset.
Spell checking: Bloom filters can be used to store a dictionary of correctly spelled words, allowing for fast spell checking by querying the filter for a given word.