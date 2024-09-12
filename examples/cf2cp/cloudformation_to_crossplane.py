import yaml
import json
import argparse
from jinja2 import Template

# Load CloudFormation template (YAML or JSON)
def load_cloudformation_template(file_path):
    with open(file_path, 'r') as stream:
        try:
            if file_path.endswith(".json"):
                return json.load(stream)
            else:
                return yaml.safe_load(stream)
        except (yaml.YAMLError, json.JSONDecodeError) as exc:
            print(exc)
            return None

# Translate CloudFormation to Crossplane manifest
def cloudformation_to_crossplane(resource_type, resource_props):
    # Mapping CloudFormation resources to Crossplane CRDs

    # Check for S3 Bucket
    if resource_type == 'AWS::S3::Bucket':
        bucket_name = resource_props.get('BucketName', 'example-bucket')
        region = resource_props.get('Region', 'us-east-1')  # Default region, update as needed

        # Crossplane manifest for S3 bucket
        crossplane_template = '''
apiVersion: s3.aws.upbound.io/v1beta1
kind: Bucket
metadata:
  name: {{ bucket_name }}
spec:
  forProvider:
    region: {{ region }}
'''
        template = Template(crossplane_template)
        crossplane_output = template.render(bucket_name=bucket_name, region=region)
        return crossplane_output

    # Check for EC2 Instance
    elif resource_type == 'AWS::EC2::Instance':
        region = resource_props.get('Region', 'us-east-1')  # Default region, update as needed
        instance_type = resource_props.get('InstanceType', 't3.micro')
        ami = resource_props.get('ImageId', 'ami-12345678')  # Default AMI

        # Crossplane manifest for EC2 instance
        crossplane_template = '''
apiVersion: ec2.aws.upbound.io/v1beta1
kind: Instance
metadata:
  name: {{ instance_name }}
spec:
  forProvider:
    region: {{ region }}
    instanceType: {{ instance_type }}
    ami: {{ ami }}
'''
        template = Template(crossplane_template)
        crossplane_output = template.render(instance_name='example-instance', region=region,instance_type=instance_type, ami=ami)
        return crossplane_output

    # Check for RDS Cluster
    elif resource_type == 'AWS::RDS::DBCluster':
        db_cluster_identifier = resource_props.get('DBClusterIdentifier', 'example-cluster')
        engine = resource_props.get('Engine', 'aurora')
        master_username = resource_props.get('MasterUsername', 'admin')

        # Crossplane manifest for RDS cluster
        crossplane_template = '''
apiVersion: rds.aws.upbound.io/v1beta1
kind: Cluster
metadata:
  name: {{ db_cluster_identifier }}
spec:
  forProvider:
    engine: {{ engine }}
    masterUsername: {{ master_username }}
'''
        template = Template(crossplane_template)
        crossplane_output = template.render(db_cluster_identifier=db_cluster_identifier, engine=engine, master_username=master_username)
        return crossplane_output

    else:
        raise ValueError("Unsupported CloudFormation resource type: " + resource_type)

# Save generated Crossplane manifest to file
def save_crossplane_file(output, file_path):
    with open(file_path, 'w') as f:
        f.write(output)
        print(f"Crossplane manifest saved to {file_path}")

if __name__ == "__main__":
    # Set up command-line argument parsing
    parser = argparse.ArgumentParser(description="Convert CloudFormation template to Crossplane manifest")
    parser.add_argument('--input', '-i', required=True, help='Path to the CloudFormation template (JSON or YAML)')
    parser.add_argument('--output', '-o', required=True, help='Path to save the generated Crossplane manifest')

    # Parse the command-line arguments
    args = parser.parse_args()

    # Load CloudFormation template
    cloudformation_template = load_cloudformation_template(args.input)

    # Check if the template was loaded correctly
    if cloudformation_template is None:
        print("Error: Could not load the CloudFormation template.")
        exit(1)

    # Get resources from CloudFormation template
    resources = cloudformation_template.get('Resources', {})
    
    # Convert each resource in the CloudFormation template to Crossplane
    crossplane_manifests = []
    for resource_name, resource_details in resources.items():
        resource_type = resource_details.get('Type')
        resource_props = resource_details.get('Properties', {})

        try:
            crossplane_manifest = cloudformation_to_crossplane(resource_type, resource_props)
            crossplane_manifests.append(crossplane_manifest)
        except ValueError as e:
            print(e)

    # Combine and save all generated Crossplane manifests
    if crossplane_manifests:
        output = "\n---\n".join(crossplane_manifests)
        save_crossplane_file(output, args.output)
    else:
        print("No resources were converted.")
