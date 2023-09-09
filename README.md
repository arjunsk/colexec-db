### Vectorized (Columnar) Execution Engine


#### Introduction
- This is a very primitive Vectorized Query Execution Engine extracted 
from [MatrixOne](https://github.com/matrixorigin/matrixone) database.

- Some of the packages have be renamed and ordered for easy understanding.

- This supports a very basic Projection Relational Algebra and Supports Abs Function. 
This should hopefully give you a high level overview of how Columnar Execution Works.

#### Pending issues
- Passing `Attr` to the result `Batch`
- Documentation
- Implement a Parquet based Storage Engine/Reader