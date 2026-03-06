$T(n)$ is the worst case running time function.
## $\Theta$ - Notation

For a given function $g(n)$, we denote by $\Theta(g(n))$ the set of functions 
```math
\Theta(g(n)) = \{f(n) : \text{there exists positive constants } c_1, c_2, \text{ and } n_0 \text{ such that } 0 \leq c_1g(n) \leq f(n) \leq c_2g(n) \text{ for all } n \geq n_0\}
```
Because $\Theta(g(n))$ is a set of functions, we can say $f(n) \in \Theta(g(n))$. 
However, it is common to simply say $f(n) = \Theta(g(n))$ as it has its advantages.
We say that $g(n)$ is an asymptotically tight bound for $f(n)$.

Example

Claim:
```math
\frac{1}{2}n^2-3n = \Theta(n^2)
```
Proof:

Given
```math
f(n) = \frac{1}{2}n^2-3n, g(n) = n^2
```

```math
0 \leq c_1n^2 \leq \frac{1}{2}n^2-3n \leq c_2n^2
```
Divide by $n^2$
```math
0 \leq c_1\leq \frac{1}{2}-\frac{3}{n} \leq c_2
```
Let
```math 
c_1 = \frac{1}{4},   \text{ and } c_2 = \frac{1}{2}
```
```math
0 \leq \frac{1}{4}\leq \frac{1}{2}-\frac{3}{n} \leq \frac{1}{2}
```
Solve for n
```math
\frac{1}{4} \leq \frac{1}{2} - \frac{3}{n}
```
```math
-\frac{1}{4}\leq- \frac{3}{n}
```
```math
1 \geq \frac{12}{n}
```
```math
n \geq 12
```
Therefore let
```math
n_0 = 12
```
Finally
```math
0 \leq \frac{1}{4}\leq \frac{1}{2}-\frac{3}{12} \leq \frac{1}{2} 
```

We can even show any quadratic function $f(n) = an^2 + bn + c \in \Theta (n^2)$. 
We can even push this to show that any polynomial $p(n) = \sum_{i=0}^{d}a_in^i$ where $a_i$ are constants and $a_d > 0$.
Then $p(n) \in \Theta(n^d)$. 

This is the reason the lower-order terms and the leading coefficient of the highest-order term is dropped when dealing with asymptotic notation.

## $O$ - Notation
## $\Omega$ - Notation
## $o$ - Notation
## $\omega$ - Notation