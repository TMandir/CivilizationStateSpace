# Contributing

Contributions to the experimental Civilization State model are welcome for
review.

## Before contributing

- Keep changes focused on the simplified structure of a civilization state
  and its future development.
- Preserve existing exported type and field names unless a change is
  explicitly coordinated; downstream code may depend on them.
- Do not imply that the current placeholders implement quantum behavior,
  operational rules, or other functionality that is not present.
- Discuss larger structural changes before submitting them.

## Changes and tests

1. Make the smallest change that fully addresses the proposed improvement.
2. Add or update tests for changed behavior or structure.
3. Format Go files with `gofmt`.
4. Run the package test suite:

   ```sh
   go test ./...
   ```

5. In your contribution description, summarize the change and include the
   test command and its result.

## License

This repository is All Rights Reserved. Contributions do not grant permission
to use, copy, modify, distribute, sublicense, or sell the software. See
[LICENSE](./LICENSE).
