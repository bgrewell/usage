# update_readme.py
import sys

def update_readme(help_output_file):
    with open('README.md', 'r') as f:
        readme = f.readlines()

    with open(help_output_file, 'r') as f:
        help_output = f.readlines()

    # Find the start and end indices of the help output section in README.md
    start_index = readme.index('```text\n')
    end_index = readme.index('```\n', start_index)

    # Replace the help output section with the new help output
    readme[start_index+1:end_index] = help_output

    # Write the updated README.md file
    with open('README.md', 'w') as f:
        f.writelines(readme)

if __name__ == "__main__":
    help_output_file = sys.argv[1]
    update_readme(help_output_file)
