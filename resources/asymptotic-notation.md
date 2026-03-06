$T(n)$ is the worst case running time function.
## $\Theta$ - Notation
The $\Theta$-notation asymptotically bounds a function from above and below.

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
When we only have an asymptotic upper bound we use $O$-notation.

For a given function $g(n)$, we denote by $O(g(n))$ the set of functions 
```math
O(g(n)) = \{f(n) : \text{there exists positive constants } c \text{ and } n_0 \text{ such that } 0 \leq  f(n) \leq cg(n) \text{ for all } n \geq n_0\}
```
$O$-notation describes an asymptotic upper bound not an asymptotically tight upper bound. Therefore we can say n = $O(n^2)$. 

When we say the running time of an alogirthm is $O(g(n))$ we usually mean "The worst case running time is $O(g(n))$."



## $\Omega$ - Notation
$\Omega$-Notation provides an asymptotic lower bound. 

For a given function $g(n)$, we denote by $\Omega(g(n))$ the set of functions 
```math
\Omega(g(n)) = \{f(n) : \text{there exists positive constants } c \text{ and } n_0 \text{ such that } 0 \leq cg(n) \leq f(n) \text{ for all } n \geq n_0\}
```
### Theorem 
Given two functions $f(n)$ and $g(n)$, 

```math
f(n) \in \Theta(g(n)) \iff f(n)\in O(g(n)) \cap f(n)\in \Omega(g(n))
```
We usually use this theorem to prove asymptotically tight bounds on a function. 

When we say the running time of an algorithm is $\Omega(g(n))$ we usually mean "The best case running time is $\Omega(g(n))$."

## $o$ - Notation
The main difference between $O$-notation and $o$-notation is that $o$-notation provides a strictly upper bound on a function while $O$-notation could allow a tight upper bound on a function.

For a given function $g(n)$, we denote by $o(g(n))$ the set of functions 
```math
o(g(n)) = \{f(n) : \text{for all positive constants } c > 0\text{, there exists a constant } n_0 > 0 \text{ such that } 0 \leq  f(n) < cg(n) \text{ for all } n \geq n_0\}
```

Examples

```math
2n \in O(n^2) 
```
```math
2n \in O(n)
```
```math
2n \in o(n^2) 
```
```math
2n \notin o(n)
```

## $\omega$ - Notation
Similarly to $o$-notation, the main difference between $\Omega$-notation and $\omega$-notation is that $\omega$-notation provides a strictly lower bound on a function while $\Omega$-notation could allow a tight lower bound on a function.

For a given function $g(n)$, we denote by $\omega(g(n))$ the set of functions 
```math
\omega(g(n)) = \{f(n) : \text{for all positive constants } c > 0\text{, there exists a constant } n_0 > 0 \text{ such that } 0 \leq c(g(n)) < f(n) \text{ for all } n \geq n_0\}
```

Examples

```math
2n \in \Omega(n) 
```
```math
2n \in \Omega(1)
```
```math
2n \notin \omega(n) 
```
```math
2n \in \omega(1)
```

## Source

These are my personal study notes and summarize material from:

Cormen, T. H., Leiserson, C. E., Rivest, R. L., & Stein, C.  
*Introduction to Algorithms*, 3rd Edition, MIT Press.  
Section 3.1 – Asymptotic Notation.