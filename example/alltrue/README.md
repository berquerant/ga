# alltrue

- Expect that gene's sections become all true, like 1111.
- Optimize the gene by repeating 100 generations.
- Each of generations has 30 individuals.
- The length of a gene is 32.
- 90% chance of crossover. It is the uniform crossover. The targets are selected by the score roulette.
- 8% chance of regeneration of a gene, just copy.
- 2% chance of the mutation which reverses 1 section of a gene.

The evaluation function is:

```
eval(gene) = encourage(count(gene))
where
  count(gene) = len(x for x in gene if x)
  encourage(score) = (score + (score - 20) ^ 2 / 2) if score > 20 else score
```
